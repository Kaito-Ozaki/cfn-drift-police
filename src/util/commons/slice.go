package comutil

/**
DeleteByListは、対象となるList（in）から、削除対象List（deleteTargetList）に含まれる要素を全て削除する関数です。
*/
func DeleteByList(in []string, deleteTargetList []string) []string {
	for _, v := range deleteTargetList {
		in = Delete(in, v)
	}
	return in
}

/**
Deleteは、Listの要素のうち、削除対象と同一のものを削除する関数です。
*/
func Delete(in []string, deleteTarget string) []string {
	out := []string{}
	for _, v := range in {
		if v != deleteTarget {
			out = append(out, v)
		}
	}
	return out
}
