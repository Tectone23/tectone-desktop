package genesis

import (
	"context"
	"encoding/json"
	"os"
	"tectone/tectone-desktop/internal/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const DefaultDevNetGenesis = `{
  "alloc": [
    {
      "addr": "7777777777777777777777777777777777777777777777777774MSJUVU",
      "comment": "RewardsPool",
      "state": {
        "algo": 125000000000000,
        "onl": 2
      }
    },
    {
      "addr": "A7NMWS3NT3IUDMLVO26ULGXGIIOUQ3ND2TXSER6EBGRZNOBOUIQXHIBGDE",
      "comment": "FeeSink",
      "state": {
        "algo": 100000,
        "onl": 2
      }
    },
    {
      "addr": "ZQ45INL2IMX66LP4ZEI2WNPESMGQKKVQADISKXNXRJKNXKJFKANOVMRYCQ",
      "comment": "Wallet1",
      "state": {
        "algo": 300000000000000,
        "onl": 1,
        "sel": "jx49tNyOGVz/yKkmjmJIFzta5SyeJxyYD7znjiwwlRs=",
        "vote": "CG6x+hWIFGJcSZVFuuRvtDD0CF1SCpeO4Hxe4YTVojQ=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "FCEV42VP5ESVQY7NZT2B327NXE53HHND5TCIXYRIANBEH2PBOD3W4EQPTM",
      "comment": "Wallet10",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "HexL5RAyE37UC/fd1UH/6Kwbomit6MkLtpe1Pi7USbU=",
        "vote": "bTqiNA6+GWmeLQjdiuKvMpiBEFfh/SGRTUewF/BPzUY=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "4HT4YQU2WPVV6EYPWYJ3L3IF2MOXQEBMJJOWMGVCBQY2TFFOB5QXMV5EPY",
      "comment": "Wallet11",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "hP67tzuIoZrXYO9Ln4I8ce/jYtfhpEaqAFoP0Gk4mSk=",
        "vote": "lA0se5QSffyxIjKlMntM/0ge5QtDaLGnwI/0GZC3BfU=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "VKAXPXHZQX5Z6ECNEDV6RFKM3QANXTGX23Z7O3PPRLCTQGZ4WHVCYJB5X4",
      "comment": "Wallet12",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "N5ivQdIOTRHAuaazCPwqGB6R0vzs4oE5C/HX5QoWJ1w=",
        "vote": "c2KiifMNoLuCpEnvLL97c6b9YuRdqhKntXyQizV7rQA=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "HWNVDWBYNO42ZEH46UVHV26HY5ABWZ2EGY5RY4SZY5SGPGMBCWPVP4GRHU",
      "comment": "Wallet13",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "lBU/qgAYZFI4HOWNA4ENWjNyS2QMNdcuh/tgeJgfuBY=",
        "vote": "+3YbNs94emm4vUzQm+rIYcKwdF5PuHcnKL5ork5pUZg=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "AK3JQTFREJMQQXKYY7DGDGMXVRDUKK63DKNXJG3BTP6EMLZ62JD455Y5TM",
      "comment": "Wallet14",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "zIJa2IwZb5wBeXgEDN7oI9Z8mgBXySRJYulW/3klDqQ=",
        "vote": "bOLmPR2wdEohQfA3OzEr5SuJyZfsyCQoOS4kxBIhO8s=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "75DBTPH3U4NBJ5IYLVEFSBVN3PIRF3OGBKTSQGJFLQBFN6MYXPZLDTQ4QI",
      "comment": "Wallet15",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "gjGwMoylS7mmmn0xcO5s91X6jyqMvyfOQykBTUqdSZo=",
        "vote": "SeNznA406oiGF/8CL/SUNZuQ9Er6o8syd3giX34jRmo=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "5K7GCWCI4MT2LWAWFP4WZUHU4WNB6AWZ5NEJCX2DSIRBEYECOW6WJSHVME",
      "comment": "Wallet16",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "C6eOqo65UBvVkzp8UpO7Wp/v89rfIChsHR/2C5nPcSE=",
        "vote": "87ZcLZx04HFyf6XUKkm9inOdYmCrnXyWw84BMBBwtUc=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "ZJMREOYU3USMDCQDMWVRECY6JGZD6LC44JTY4TC55ZYTXTQUBFRBEIPY6A",
      "comment": "Wallet17",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "l77BpNvhHEEiAjY+RNi8ht62+qgb90DXNqKzB+4uHew=",
        "vote": "9iuQKJNBXjiKTRhPHPYqEJlmLo0g04in4ABdo7elkFA=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "XLI4GVZKYHXVK4MU2YRQMT6TZVCB772NLWW7TEGQIGCH4JXSZA6PARNBVQ",
      "comment": "Wallet18",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "gAweEod0UC6CN1fjFt/ZFVbyw2kON3hcG5HSPUY7X7A=",
        "vote": "1/Vyqjf5sHiZrq5StV+lOK8I0NZCSUbQ5p4QiCFW4a0=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "XLQBCY6TK5JSDOIMC4U22ZQ7O35VHJS6QWQUQY6KFV2HD4T3QXCNOEZHMY",
      "comment": "Wallet19",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "P3zYirarJ4QIMGz2h4SK6rCD5wJMrGPtfIop99ZTCYg=",
        "vote": "J3kmHUeU21FvOP46A9OB1QUamatT32oBy2dmZQFVfi0=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "LVFCXGJ4H2QXLXS67H7JPQ6HDLSIQ4UBDJ7RNWKI6UUG4QC3GZX53R3BLA",
      "comment": "Wallet2",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "VyZiFcfly5SCIZhs79IbxYbfYw8od38ecs6vm8UH760=",
        "vote": "hGNbwxyqnTEq8rw3d6Ezt4kYDVsk49X4Ousi6z2MPaU=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "2BMHSRKHDYBJTW673MLERUX24MOGSXOPFWYLMAVW3QCFRNRB4M6ZDNAFCI",
      "comment": "Wallet20",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "jHCe2j1mpV9GWRqdhHt3NtZpppVbkc1o5jUaykpdDkw=",
        "vote": "nw8oxQjT5IS8GKBWucYtKQB3kY3ywzgMjyY0+6Xi5+s=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "3LGSIMLC44UB4NQHSFVM6ZDE7BVJFJVDVYESZZ4SJU4EYDZQV32CIH6JM4",
      "comment": "Wallet3",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "jGQIdnRvYaokSw7iRG29AdSR81OFqgdnU3sFvd/g0gI=",
        "vote": "3bu4xBXbFh/Bpqtdwg9eqlUXGa6nkVE6QDf2gSr7kng=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "UXECRB7D5NDQEKVK5ISA5FUKPX6UC3XFI5JEJEYV3CUHJRMEYY2ZBUGVHI",
      "comment": "Wallet4",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "F2GF2fUODTFsk6NyGAOYDAsS2pdeqxJ0XgJT4C+xh50=",
        "vote": "SLw4LA5EdOiUHhQ2r99fsKMytbJCQZWaUxuDrgigyLo=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "IXPBHLMKILBRSG3ZTM34UFORUQ2XHTZMEXK5LODNZTS3M3IPKPE6ZOQADY",
      "comment": "Wallet5",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "KqUKfEEC7pkR+4K4Tnp+V+Rwu/yXAbuAyXlU83ZGPNA=",
        "vote": "7vyFvIGRixNbisSJW23yfMW2wE4xRRm41IlgzaZDgjQ=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "H5AEMBGBAVVEZOLGNX22OMNDJLEU7PG5VSCBNHB46PKRBBKQBZWA7R2DNA",
      "comment": "Wallet6",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "kthetjPfSMTOa04Oo3ZQ7GGfXNNE4nCNC8tYVW9qFNE=",
        "vote": "WQBnwqMm6f9+GKR0Sz+X6ej65KCkB2I971dYtuoG7tk=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "436JPU37D2J5EGLMQCVKU236ZAMJM3XUXGGK35MZ43P3LAB3AWBZS73HME",
      "comment": "Wallet7",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "o8RSgaF1leM7Mdux9hk/cxypO4XMEZRZjrw6R3JqKZc=",
        "vote": "e44SozE14TV9B3Raam7B2Fm9rxRH9O4ruRJmdErr55w=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "TP5S4EO7CBRZMXSNMNEEJ7232DCRXHK5YN7W7SKA2W4A2Y5JBFFWCKJJKU",
      "comment": "Wallet8",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "r/wqmzTQaVYnPBZPGL8W+5i/FYqNz0/1FqewGAyRiiw=",
        "vote": "XxrwQh/SE72Q/IA18YNEUAnujR3V+ntd94S554HyQBw=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "DMOZSZ3TE65RR5UX2KLDD56LGI75W6LEVOD275S56W2AYCGDD4UYGU4URM",
      "comment": "Wallet9",
      "state": {
        "algo": 500000000000000,
        "onl": 1,
        "sel": "+LIJ3xi3SuoSuAK/RKirPUMFNltJucuFMa9H9KxB5vQ=",
        "vote": "+irtKnZKr1d4G7Fqvc7i4QYqN/NEXunPzdneHKboMUU=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "RGKUXFHCDDGN7EJHSDF533KVTCS5EFHAR57SCUV3RDOTQKIXAFRB6ME25M",
      "comment": "bank",
      "state": {
        "algo": 100000000000000,
        "onl": 1,
        "sel": "SKZZJvgPQImzzn9Ft5e7ZWitXR/1TZwk2aaBM35YcIc=",
        "vote": "jB8AMuWZn1m5go/dIVzB6DFIV8hD7UpcLLV2XGjK8FM=",
        "voteKD": 10000,
        "voteLst": 10000000
      }
    },
    {
      "addr": "VF4LPDP2OCQIN4Q2PJ4VZV725NBHMGUZZC35SO3EHX7XUB7SRMLFNRHVB4",
      "comment": "pp1",
      "state": {
        "algo": 50000000000000
      }
    },
    {
      "addr": "ZJOL26NAJMBNMAOKKTBRZTFAPSODASPKV4Z2A4EYMNMTZVT65CEAIBFTNQ",
      "comment": "pp2",
      "state": {
        "algo": 50000000000000
      }
    }
  ],
  "fees": "A7NMWS3NT3IUDMLVO26ULGXGIIOUQ3ND2TXSER6EBGRZNOBOUIQXHIBGDE",
  "id": "v1.0",
  "network": "devnet",
  "proto": "https://github.com/algorandfoundation/specs/tree/3a83c4c743f8b17adfd73944b4319c25722a6782",
  "rwd": "7777777777777777777777777777777777777777777777777774MSJUVU"
}`

func GetDefaultDevNetGenesis(ctx context.Context) model.Genesis {
	var g model.Genesis
	if err := json.Unmarshal([]byte(DefaultDevNetGenesis), &g); err != nil {
		runtime.LogErrorf(ctx, "error could not unmarshal default devnet genesis json: %s\n", err)
		return g
	}
	return g
}

func LoadDevNetGenesis(ctx context.Context, path string) model.Genesis {
	var g model.Genesis
	f, err := os.Open(path)
	if err != nil {
		runtime.LogErrorf(ctx, "could not open devnet genesis file: %s\n", err)
		return g
	}
	defer f.Close()

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		runtime.LogErrorf(ctx, "could not not read devnet genesis from file: %s\n", err)
		return g

	}

	if err := json.Unmarshal(b, &g); err != nil {
		runtime.LogErrorf(ctx, "could not unmarshal devnet genesis file: %s\n", err)
		return g
	}

	return g
}
