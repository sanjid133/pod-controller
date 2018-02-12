package ssh

import (
	"fmt"
	"testing"
)

func TestNewSSHKeyPair(t *testing.T) {
	k, err := NewSSHKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("PublicKey >>>", string(k.PublicKey))
	fmt.Println("__________________________________________________________________")
	fmt.Println("PrivateKey >>>", string(k.PrivateKey))
	fmt.Println("__________________________________________________________________")
	fmt.Println("OpensshFingerprint >>>", k.OpensshFingerprint)
	fmt.Println("__________________________________________________________________")
	fmt.Println("AwsFingerprint >>>", k.AwsFingerprint)
	fmt.Println("__________________________________________________________________")
}

/*
******************************************************
PublicKey >>> ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCxjCfrv5FTtoHCCbz+rMxFJkWQNbmPxhZij17LJZAeL8d6DauiB6K3pOOmMTWmuiyOGFnd3znk88ZG2TEu3fMgYMQ+tVmejGuE0y416IqSNFw+5HZmf/OEu/JzfF6ttlR15mSHTwdsJDUmoYpPEvCHxKcSdtddWGV2BztAmeXc5tj9aG8DyoFjfIyW/45lHgzsJ38CIRg7ZaMq9KjkG9VDfTEAB3smgFL2U/ozQfkkqEo2hI/CiDc3nYSz6mAB7cO5YXNlsVswmOF29RyImKuKgI6uAEYxtoVPDHW0mrqKY9OxsTHVGguRoBQ8mJ7j8VMmwy35xO2pk5q5li5zQ/8V
__________________________________________________________________
PrivateKey >>> -----BEGIN RSA PRIVATE KEY-----
MIIEoAIBAAKCAQEAsYwn67+RU7aBwgm8/qzMRSZFkDW5j8YWYo9eyyWQHi/Heg2r
ogeit6TjpjE1prosjhhZ3d855PPGRtkxLt3zIGDEPrVZnoxrhNMuNeiKkjRcPuR2
Zn/zhLvyc3xerbZUdeZkh08HbCQ1JqGKTxLwh8SnEnbXXVhldgc7QJnl3ObY/Whv
A8qBY3yMlv+OZR4M7Cd/AiEYO2WjKvSo5BvVQ30xAAd7JoBS9lP6M0H5JKhKNoSP
wog3N52Es+pgAe3DuWFzZbFbMJjhdvUciJirioCOrgBGMbaFTwx1tJq6imPTsbEx
1RoLkaAUPJie4/FTJsMt+cTtqZOauZYuc0P/FQIDAQABAoIBAEyKbIdgbalWK7hX
9GciXWmOJz0VPCr0LaBNyILxbaDl3iwLCqvueMKMzitOWZ/H8P9NFijlXBMM2LSg
W2il/HM/5DpYegs05AE9/VT2LA79jar7WOO1U2kskUgiHLrsuJqfOUVv13ZSYauG
xpZyRslxCMmX/B5cxljMo0mlTwgcsO7iBImDjN61yuRV/sUfF6pqj469DFt/jC50
U75pplCUCarV2TBW0Bw8Ha41RfyXkev5nJHOmCGPi74xkOlz3LipL4Q1swY7xe8M
2dwOu7qCIbVX5XbHmGwkhB2Qe3LPdIqbm0GVQHpsQCpz0ZsEyBwNTC4UGTkU2PA9
B7+5AX0CgYEA4EQfs7rtLawGJgwMA5uhokYZt0ffUHANXFxp15jf0IQMeAZxrRaY
ZRKKAQdriOsOw5SoY5w8uXIEgGveu2a8YtasddosagZK2bo0iNK9shhwhSVCPMdi
n70oIrbqPtA5eYuBhF/sM0izMSwCb3lli1gbA5wb0ZDOw7X5yeWjgd8CgYEAyquv
G5ttbe/Q2dVlvJzDIWUhO/7TMwavou4InAIzCFMoimAdVsN1W6wZHNwJuMcLst9C
8j8AsYSK84oBs784rhbofc8H3wRHFobE09FSqbKYo4DO+Tec4+VfDYmtGnto8Tac
yoqxvrxfsYD/VA9vLtBklik2jIq8wd/Bwqmw5YsCgYA2p+OiBdhv3RDSjPUQnAlU
Zwk1ZcrC68wVyl/xL0y8oLH5PnjN2kfhuhbwEULf54l9mOlA/X6/I059jcyGl48p
92+oeTeBz27GiV7sSY/5q4Mgr+fosB8VFyCE9Vc5X73iJWG3RTlndjxVq4gcatWY
lCxhQNZ2yHfbsqnhzpl+uQKBgDSxRFqqI15U3nj0OXo7S6RQFb7ydtiqYVm8B3MA
Q2JjtIOdj6MyOD34VaMiNkjbP1hnIZG3/acprbK5681e5w6kGBg1jndWNmFN8rib
kKFcNME1yRoUcm7yAK+VXvMmqNcl+sfSHXpYoYyV3ExvotmKBszftDYAEClb0pWY
5swXAn8mF9WG52j5YqTsbXQsu6ezNZiMMsGSotzSTPj9bVNjFn8ZRLq9G26+wxPp
pwIBIGp9FjALcDvzj2v1FY0hyXI2uMa86XmqMFE+kxAksWk5gW6MQSv66Aw9NJcN
NpGZaJ9b0B633if+Us7uxE/80dLN/HcsFXNadJM8Wqq2xxng
-----END RSA PRIVATE KEY-----

__________________________________________________________________
OpensshFingerprint >>> 43:54:e3:ec:50:e6:d6:e0:be:45:96:5c:51:7f:49:f9
__________________________________________________________________
AwsFingerprint >>> 65:b7:40:f0:15:99:e2:1a:bb:84:f7:e6:28:26:e5:3f
__________________________________________________________________
*/

func TestSSHKeyPairFormat(t *testing.T) {
	pub := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCxjCfrv5FTtoHCCbz+rMxFJkWQNbmPxhZij17LJZAeL8d6DauiB6K3pOOmMTWmuiyOGFnd3znk88ZG2TEu3fMgYMQ+tVmejGuE0y416IqSNFw+5HZmf/OEu/JzfF6ttlR15mSHTwdsJDUmoYpPEvCHxKcSdtddWGV2BztAmeXc5tj9aG8DyoFjfIyW/45lHgzsJ38CIRg7ZaMq9KjkG9VDfTEAB3smgFL2U/ozQfkkqEo2hI/CiDc3nYSz6mAB7cO5YXNlsVswmOF29RyImKuKgI6uAEYxtoVPDHW0mrqKY9OxsTHVGguRoBQ8mJ7j8VMmwy35xO2pk5q5li5zQ/8V`
	priv := `-----BEGIN RSA PRIVATE KEY-----
MIIEoAIBAAKCAQEAsYwn67+RU7aBwgm8/qzMRSZFkDW5j8YWYo9eyyWQHi/Heg2r
ogeit6TjpjE1prosjhhZ3d855PPGRtkxLt3zIGDEPrVZnoxrhNMuNeiKkjRcPuR2
Zn/zhLvyc3xerbZUdeZkh08HbCQ1JqGKTxLwh8SnEnbXXVhldgc7QJnl3ObY/Whv
A8qBY3yMlv+OZR4M7Cd/AiEYO2WjKvSo5BvVQ30xAAd7JoBS9lP6M0H5JKhKNoSP
wog3N52Es+pgAe3DuWFzZbFbMJjhdvUciJirioCOrgBGMbaFTwx1tJq6imPTsbEx
1RoLkaAUPJie4/FTJsMt+cTtqZOauZYuc0P/FQIDAQABAoIBAEyKbIdgbalWK7hX
9GciXWmOJz0VPCr0LaBNyILxbaDl3iwLCqvueMKMzitOWZ/H8P9NFijlXBMM2LSg
W2il/HM/5DpYegs05AE9/VT2LA79jar7WOO1U2kskUgiHLrsuJqfOUVv13ZSYauG
xpZyRslxCMmX/B5cxljMo0mlTwgcsO7iBImDjN61yuRV/sUfF6pqj469DFt/jC50
U75pplCUCarV2TBW0Bw8Ha41RfyXkev5nJHOmCGPi74xkOlz3LipL4Q1swY7xe8M
2dwOu7qCIbVX5XbHmGwkhB2Qe3LPdIqbm0GVQHpsQCpz0ZsEyBwNTC4UGTkU2PA9
B7+5AX0CgYEA4EQfs7rtLawGJgwMA5uhokYZt0ffUHANXFxp15jf0IQMeAZxrRaY
ZRKKAQdriOsOw5SoY5w8uXIEgGveu2a8YtasddosagZK2bo0iNK9shhwhSVCPMdi
n70oIrbqPtA5eYuBhF/sM0izMSwCb3lli1gbA5wb0ZDOw7X5yeWjgd8CgYEAyquv
G5ttbe/Q2dVlvJzDIWUhO/7TMwavou4InAIzCFMoimAdVsN1W6wZHNwJuMcLst9C
8j8AsYSK84oBs784rhbofc8H3wRHFobE09FSqbKYo4DO+Tec4+VfDYmtGnto8Tac
yoqxvrxfsYD/VA9vLtBklik2jIq8wd/Bwqmw5YsCgYA2p+OiBdhv3RDSjPUQnAlU
Zwk1ZcrC68wVyl/xL0y8oLH5PnjN2kfhuhbwEULf54l9mOlA/X6/I059jcyGl48p
92+oeTeBz27GiV7sSY/5q4Mgr+fosB8VFyCE9Vc5X73iJWG3RTlndjxVq4gcatWY
lCxhQNZ2yHfbsqnhzpl+uQKBgDSxRFqqI15U3nj0OXo7S6RQFb7ydtiqYVm8B3MA
Q2JjtIOdj6MyOD34VaMiNkjbP1hnIZG3/acprbK5681e5w6kGBg1jndWNmFN8rib
kKFcNME1yRoUcm7yAK+VXvMmqNcl+sfSHXpYoYyV3ExvotmKBszftDYAEClb0pWY
5swXAn8mF9WG52j5YqTsbXQsu6ezNZiMMsGSotzSTPj9bVNjFn8ZRLq9G26+wxPp
pwIBIGp9FjALcDvzj2v1FY0hyXI2uMa86XmqMFE+kxAksWk5gW6MQSv66Aw9NJcN
NpGZaJ9b0B633if+Us7uxE/80dLN/HcsFXNadJM8Wqq2xxng
-----END RSA PRIVATE KEY-----
`

	k, err := ParseSSHKeyPair(pub, priv)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(k)
}