package comutil

/**
StringCtoPは、string型の定数をポインタに変換する関数です。
正確には、与えられた定数と同じ値を持つ変数を作成し、そのポインタを返却しています。
*/
func StringCtoP(con string) *string {
	return &con
}

/**
Int64CtoPは、int64型の定数をポインタに変換する関数です。
正確には、与えられた定数と同じ値を持つ変数を作成し、そのポインタを返却しています。
*/
func Int64CtoP(con int64) *int64 {
	return &con
}

/**
BoolCtoPは、bool型の定数をポインタに変換する関数です。
正確には、与えられた定数と同じ値を持つ変数を作成し、そのポインタを返却しています。
*/
func BoolCtoP(con bool) *bool {
	return &con
}

/**
StringSclieToStringPSliceは、Stringを要素とするSliceを、Stringポインタを要素とするSliceに変換する関数です。
*/
func StringSliceToStringPSlice(ss []string) []*string {
	stringSliceToP := []*string{}
	for i := range ss {
		stringSliceToP = append(stringSliceToP, &ss[i])
	}
	return stringSliceToP
}

/**
StringSclieToStringPSliceは、Stringポインタを要素とするSliceを、Stringを要素とするSliceに変換する関数です。
*/
func StringPSliceToStringSlice(sps []*string) []string {
	stringSlice := []string{}
	for _, s := range sps {
		stringSlice = append(stringSlice, *s)
	}
	return stringSlice
}
