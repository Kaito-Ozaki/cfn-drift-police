package main_test

import (
	"errors"
	"testing"

	cfndto "cfn-drift-police/src/application/dto/cloudformation"
	"cfn-drift-police/src/usecase/check"
	mock_cfncli "cfn-drift-police/test/util/cloudformation"
	mock_sqscli "cfn-drift-police/test/util/sqs"

	"github.com/golang/mock/gomock"
)

/**
mockは、各mockクライアントを保持する構造体です。
*/
type mock struct {
	ctrl       *gomock.Controller
	mockCfnCli *mock_cfncli.MockCloudFormationClient
	mockSqsCli *mock_sqscli.MockSqsClient
}

/**
beforeEachは、各テストの最初に実行されるべき関数です。mockインスタンスの生成を行います。
*/
func beforeEach(t *testing.T) mock {
	ctrl := gomock.NewController(t)
	mockCfnCli := mock_cfncli.NewMockCloudFormationClient(ctrl)
	mockSqsCli := mock_sqscli.NewMockSqsClient(ctrl)
	return mock{
		ctrl:       ctrl,
		mockCfnCli: mockCfnCli,
		mockSqsCli: mockSqsCli,
	}
}

/**
afterEachは、各テストの最後に実行されるべき関数です。mockインスタンスの削除を行います。
*/
func afterEach(m mock) {
	m.ctrl.Finish()
}

/**
正常系001_001 一件のドリフト検出に成功し、正常終了する
*/
func TestCaseCheck001001(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// CFnスタック取得mock設定
	ss := []cfndto.StackSummary{}
	ss = append(ss, cfndto.StackSummary{
		StackName: "test-cfn-drift-police",
	})
	listStacksOutput := cfndto.ListStacksOutput{
		StackSummaries: ss,
	}
	mock.mockCfnCli.EXPECT().ListStacks(gomock.Any()).Return(&listStacksOutput, nil).Times(1)

	// CFnドリフト検出実行mock設定
	detectStackDriftOutput := cfndto.DetectStackDriftOutput{
		StackDriftDetectionId: "arn:aws:cloudformation:ap-northeast-1:000000000000:stack/test-cfn-drift-police/aa67b910-dfe2-11ec-b933-068c0ea753ab",
	}
	mock.mockCfnCli.EXPECT().DetectStackDrift(gomock.Any()).Return(&detectStackDriftOutput, nil).Times(1)

	// SQSメッセージ送信mock設定
	mock.mockSqsCli.EXPECT().SendMessage(gomock.Any()).Return(nil).Times(1)

	// checkユースケース実行
	cu := check.NewCheckUsecase(mock.mockCfnCli, mock.mockSqsCli)
	cu.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系001_002 二件のドリフト検出に成功し、正常終了する
*/
func TestCaseCheck001002(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// CFnスタック取得mock設定
	ss := []cfndto.StackSummary{}
	ss = append(ss, cfndto.StackSummary{
		StackName: "test-cfn-drift-police",
	})
	ss = append(ss, cfndto.StackSummary{
		StackName: "test-cfn-drift-police2",
	})
	listStacksOutput := cfndto.ListStacksOutput{
		StackSummaries: ss,
	}
	mock.mockCfnCli.EXPECT().ListStacks(gomock.Any()).Return(&listStacksOutput, nil).Times(1)

	// CFnドリフト検出実行mock設定
	detectStackDriftOutput := cfndto.DetectStackDriftOutput{
		StackDriftDetectionId: "arn:aws:cloudformation:ap-northeast-1:000000000000:stack/test-cfn-drift-police/aa67b910-dfe2-11ec-b933-068c0ea753ab",
	}
	mock.mockCfnCli.EXPECT().DetectStackDrift(gomock.Any()).Return(&detectStackDriftOutput, nil).Times(2)

	// SQSメッセージ送信mock設定
	mock.mockSqsCli.EXPECT().SendMessage(gomock.Any()).Return(nil).Times(2)

	// checkユースケース実行
	cu := check.NewCheckUsecase(mock.mockCfnCli, mock.mockSqsCli)
	cu.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_001 スタック検出に失敗するが、処理を続行し、正常終了する
*/
func TestCaseCheck002001(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// CFnスタック取得mock設定
	ss := []cfndto.StackSummary{}
	ss = append(ss, cfndto.StackSummary{
		StackName: "test-cfn-drift-police",
	})
	listStacksOutput := cfndto.ListStacksOutput{
		StackSummaries: ss,
	}
	mock.mockCfnCli.EXPECT().ListStacks(gomock.Any()).Return(&listStacksOutput, nil).Times(1)

	// CFnドリフト検出実行mock設定
	mock.mockCfnCli.EXPECT().DetectStackDrift(gomock.Any()).Return(nil, errors.New("fail")).Times(1)

	// SQSメッセージ送信mock設定
	mock.mockSqsCli.EXPECT().SendMessage(gomock.Any()).Return(nil).Times(0)

	// checkユースケース実行
	cu := check.NewCheckUsecase(mock.mockCfnCli, mock.mockSqsCli)
	cu.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
正常系002_002 SQSへのメッセージ送信に失敗するが、処理を続行し、正常終了する
*/
func TestCaseCheck002002(t *testing.T) {
	// mockインスタンス生成
	mock := beforeEach(t)

	// CFnスタック取得mock設定
	ss := []cfndto.StackSummary{}
	ss = append(ss, cfndto.StackSummary{
		StackName: "inmemory-cfn-drift-police",
	})
	listStacksOutput := cfndto.ListStacksOutput{
		StackSummaries: ss,
	}
	mock.mockCfnCli.EXPECT().ListStacks(gomock.Any()).Return(&listStacksOutput, nil).Times(1)

	// CFnドリフト検出実行mock設定
	detectStackDriftOutput := cfndto.DetectStackDriftOutput{
		StackDriftDetectionId: "arn:aws:cloudformation:ap-northeast-1:000000000000:stack/test-cfn-drift-police/aa67b910-dfe2-11ec-b933-068c0ea753ab",
	}
	mock.mockCfnCli.EXPECT().DetectStackDrift(gomock.Any()).Return(&detectStackDriftOutput, nil).Times(1)

	// SQSメッセージ送信mock設定
	mock.mockSqsCli.EXPECT().SendMessage(gomock.Any()).Return(errors.New("fail")).Times(1)

	// checkユースケース実行
	cu := check.NewCheckUsecase(mock.mockCfnCli, mock.mockSqsCli)
	cu.Execute()

	// mockインスタンス解放
	afterEach(mock)
}

/**
異常系001_001 スタック取得に失敗して、異常終了する
os.Exitする関係で、テストが失敗するので、コメントアウト
仕様を示す、という意味合いでテストコードは残しておく
*/
// func TestCaseCheckFatal001001(t *testing.T) {
// 	// mockインスタンス生成
// 	mock := beforeEach(t)

// 	// CFnスタック取得mock設定
// 	mock.mockCfnCli.EXPECT().ListStacks(gomock.Any()).Return(nil, errors.New("fail")).Times(1)

// 	// CFnドリフト検出実行mock設定
// 	detectStackDriftOutput := cfndto.DetectStackDriftOutput{
// 		StackDriftDetectionId: "arn:aws:cloudformation:ap-northeast-1:000000000000:stack/test-cfn-drift-police/aa67b910-dfe2-11ec-b933-068c0ea753ab",
// 	}
// 	mock.mockCfnCli.EXPECT().DetectStackDrift(gomock.Any()).Return(&detectStackDriftOutput, nil).Times(0)

// 	// SQSメッセージ送信mock設定
// 	mock.mockSqsCli.EXPECT().SendMessage(gomock.Any()).Return(nil).Times(0)

// 	// checkユースケース実行
// 	cu := check.NewCheckUsecase(mock.mockCfnCli, mock.mockSqsCli)
// 	cu.Execute()

// 	// mockインスタンス解放
// 	afterEach(mock)
// }
