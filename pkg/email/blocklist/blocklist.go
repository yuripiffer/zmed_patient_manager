package blocklist

func IsDomainInBlockList(emailDomain string) bool {
	//TODO add more blockList servers
	domainBlockList := []string{
		"spam.uk",
		"spam.com",
		"test.com",
	}

	for _, blockedDomain := range domainBlockList {
		if blockedDomain == emailDomain {
			return true
		}
	}
}
