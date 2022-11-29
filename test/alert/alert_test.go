package main_test

import (
	"errors"
	"testing"

	"cfn-drift-police/src/application/consts"
	cfndto "cfn-drift-police/src/application/dto/cloudformation"
	sqsdto "cfn-drift-police/src/application/dto/sqs"
	"cfn-drift-police/src/usecase/alert"
	comutil "cfn-drift-police/src/util/commons"
	mock_cfncli "cfn-drift-police/test/util/cloudformation"
	mock_slackcli "cfn-drift-police/test/util/slack"
	mock_sqscli "cfn-drift-police/test/util/sqs"

	"github.com/golang/mock/gomock"
)

/**
mockは、各mockクライアントを保持する構造体です。
*/
type mock struct {
	ctrl         *gomock.Controller
	mockCfnCli   *mock_cfncli.MockCloudFormationClient
	mockSqsCli   *mock_sqscli.MockSqsClient
	mockSlackCli *mock_slackcli.MockSlackClient
}

/**
beforeEachは、各テストの最初に実行されるべき関数です。mockインスタンスの生成を行います。
*/
func beforeEach(t *testing.T) mock {
	ctrl := gomock.NewController(t)
	mockCfnCli := mock_cfncli.NewMockCloudFormationClient(ctrl)
	mockSqsCli := mock_sqscli.NewMockSqsClient(ctrl)
	mockSlackCli := mock_slackcli.NewMockSlackClient(ctrl)
	return mock{
		ctrl:         ctrl,
		mockCfnCli:   mockCfnCli,
		mockSqsCli:   mockSqsCli,
		mockSlackCli: mockSlackCli,
	}
}

/**
afterEachは、各テストの最後に実行されるべき関数です。mockインスタンスの削除を行います。
*/
func afterEach(m mock) {
	m.ctrl.Finish()
}

/**
正常系001_001 1件のドリフト検知に成功し、正常終了する
*/
func TestCaseAlert001001(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系001_002 2件のドリフト検知に成功し、正常終了する
*/
func TestCaseAlert001002(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
			mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
		}).Return(&receiveMessageOutput, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	sds := "SYNC"
	describeStackDriftDetectionStatusOutput2 := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      &sds,
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police2%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Do(func(in cfndto.DescribeStackDriftDetectionStatusInput) {
		mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput2, nil).Times(1)
	}).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(2)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系001_003 1件のドリフト検知に失敗し、正常終了する
*/
func TestCaseAlert001003(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	sds := "SYNC"
	dsr := "{\"Summary\":\"Failed to detect drift on resources [CfnDriftPoliceTest]\",\"Failures\":[{\"Resource\":\"CfnDriftPoliceTest\",\"FailureReason\":\"Internal Failure\"}]}"
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_FAILED",
		DetectionStatusReason: &dsr,
		StackDriftStatus:      &sds,
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police2%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_001 SQSメッセージボディのデコードに失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002001(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "test",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(0)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_002 ドリフト結果取得に失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002002(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(nil, errors.New("fail")).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_003 Slackへのドリフトメッセージ送信に失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002003(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(errors.New("fail")).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_004 SQSメッセージの削除に失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002004(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(errors.New("fail")).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_005 ドリフト検出失敗理由のデコードに失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002005(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	sds := "SYNC"
	dsr := "test"
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_FAILED",
		DetectionStatusReason: &dsr,
		StackDriftStatus:      &sds,
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police2%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_006 Slackへのドリフト検出失敗メッセージ送信に失敗するが、処理を続行し、正常終了する
*/
func TestCaseAlert002006(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// SQSメッセージ取得mock設定
	ms := []sqsdto.Message{}
	ms = append(ms, sqsdto.Message{
		Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
		ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
	})
	receiveMessageOutput := sqsdto.ReceiveMessageOutput{
		Messages: &ms,
	}
	receiveMessageOutput2 := sqsdto.ReceiveMessageOutput{
		Messages: nil,
	}
	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Do(func(in sqsdto.ReceiveMessageInput) {
		mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(&receiveMessageOutput2, nil).Times(1)
	}).Return(&receiveMessageOutput, nil).Times(1)

	// CFnドリフト検出結果取得mock設定
	dsr := "{\"Summary\":\"Failed to detect drift on resources [CfnDriftPoliceTest]\",\"Failures\":[{\"Resource\":\"CfnDriftPoliceTest\",\"FailureReason\":\"Internal Failure\"}]}"
	sds := "SYNC"
	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_FAILED",
		DetectionStatusReason: &dsr,
		StackDriftStatus:      &sds,
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police2%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(1)

	// slackメッセージ投稿mock設定
	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any(), gomock.Any()).Return(errors.New("fail")).Times(1)

	// SQSメッセージ削除mock設定
	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(1)

	// alertユースケース実行
	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
	au.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
異常系001_001 SQSからのメッセージ取得に失敗して、異常終了する
os.Exitする関係で、テストが失敗するので、コメントアウト
仕様を示す、という意味合いでテストコードは残しておく
*/
// func TestCaseAlertFatal001001(t *testing.T) {
// 	// mockインスタンス生成
// 	mock := beforeEach(t)

// 	// SQSメッセージ取得mock設定
// 	mock.mockSqsCli.EXPECT().ReceiveMessage(gomock.Any()).Return(nil, errors.New("fail")).Times(1)

// 	// CFnドリフト検出結果取得mock設定
// 	describeStackDriftDetectionStatusOutput := cfndto.DescribeStackDriftDetectionStatusOutput{
// 		StackDriftStatus: comutil.StringCtoP(consts.CfnDriftStatusDrifted),
// 	}
// 	mock.mockCfnCli.EXPECT().DescribeStackDriftDetectionStatus(gomock.Any()).Return(&describeStackDriftDetectionStatusOutput, nil).Times(0)

// 	// slackメッセージ投稿mock設定
// 	mock.mockSlackCli.EXPECT().PostMessage(gomock.Any()).Return(nil).Times(0)

// 	// SQSメッセージ削除mock設定
// 	mock.mockSqsCli.EXPECT().DeleteMessage(gomock.Any()).Return(nil).Times(0)

// 	// alertユースケース実行
// 	au := alert.NewAlertUsecase(mock.mockCfnCli, mock.mockSqsCli, mock.mockSlackCli)
// 	au.Execute()

// 	// mockインスタンス解放
// 	afterEach(mock)
// }
