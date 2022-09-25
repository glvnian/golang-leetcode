package day_13

import "strconv"

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	helperRestoreIpAddresses(s, 0, 0, "", "", &result)
	return result

}

func helperRestoreIpAddresses(s string, index, seqID int, subnet, ipaddr string, result *[]string) {
	if index == len(s) && seqID == 3 && isIpNode(subnet) {
		*result = append(*result, ipaddr+subnet)
	}
	if index < len(s) && seqID <= 3 {
		ch := s[index]
		if isIpNode(subnet + string(ch)) {
			helperRestoreIpAddresses(s, index+1, seqID, subnet+string(ch), ipaddr, result)
		}

		if len(subnet) > 0 && seqID < 3 {
			helperRestoreIpAddresses(s, index+1, seqID+1, string(ch), ipaddr+subnet+".", result)
		}
	}
}

func isIpNode(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if (s[0] == '0' && len(s) > 1) || num > 255 {
		return false
	}
	return true
}
