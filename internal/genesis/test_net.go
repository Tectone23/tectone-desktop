package genesis

import (
	"context"
	"encoding/json"
	"os"
	"tectone/tectone-desktop/internal/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const DefaultTestNetGenesis = `{
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
      "addr": "LHHQJ6UMXRGEPXBVFKT7SY26BQOIK64VVPCLVRL3RNQLX5ZMBYG6ZHZMBE",
      "comment": "Wallet1",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "h7Ml/mY/PDCPSj33u72quxaMX99n+/VE+wD94/hMdzY=",
        "vote": "R9kxsHbji4DlxPOAyLehy8vaiWyLjWdLGWBLnQ5jjY8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "CQW2QBBUW5AGFDXMURQBRJN2AM3OHHQWXXI4PEJXRCVTEJ3E5VBTNRTEAE",
      "comment": "Wallet10",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "p2tiuQ2kqJGG049hHOKNIjid4/u1MqlvgXfbxK4tuEY=",
        "vote": "E73cc+KB/LGdDHO1o84440WKCmqvbM4EgROMRyHfjDc=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "HXPCXKQZF4LDL3CE5ERWC5V2BQZTKXUUT3JE6AXXNKLF3OJL4XUAW5WYXM",
      "comment": "Wallet11",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "ex32mzy8E94GkHGy+cmkRP5JNqFBKGfHtgyUGNxTiW8=",
        "vote": "BtYvtmeEBY2JovHUfePTjo3OtOMrhKp3QMeOYl3JFYM=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "Y3FSHQ43JWDSJG7LL5FBRTXHEGTPSWEQBO4CO2RO7KS2Z4ZGBUI7LSEDHQ",
      "comment": "Wallet12",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "+AtsgunCR8dzO9UGUJ6sFtAaX/E+ssK6JNmvAljQG2E=",
        "vote": "Rx21vGt6pnixU2g6NS/TknVtAGbf8hWMJiEtNuV5lb4=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "KXILJUKZJEOS4OCPGENS72JWIZOXGZSK4R235EQPGQ3JLG6R2BBT3ODXEI",
      "comment": "Wallet13",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "6s09aJVaGfPdbWy5zUSyBJEX/EGVvsn2moUOvakQdBQ=",
        "vote": "1oTW6ZpIHhQP6xeNCSqHOZZJYrKiP5D52OHXGzbVz4k=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "R4DCCBODM4L7C6CKVOV5NYDPEYS2G5L7KC7LUYPLUCKBCOIZMYJPFUDTKE",
      "comment": "Wallet14",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "XsqeQcLz5nPP316ntIp0X9OfJi5ZSfUNrlRSitWXJRg=",
        "vote": "r+e0lAD9FnNqOKoWdYdFko13pm9fk/zCJkxVVCqzjaU=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "VKM6KSCTDHEM6KGEAMSYCNEGIPFJMHDSEMIRAQLK76CJDIRMMDHKAIRMFQ",
      "comment": "Wallet15",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "64Xkj7z3rHZT7syihd0OmgNExHfnOLdLojDJZgtB1d8=",
        "vote": "um2RrGFmZ5Coned2WSbo/htYMKjW7XFE5h25M2IFsDs=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "YTOO52XR6UWNM6OUUDOGWVTNJYBWR5NJ3VCJTZUSR42JERFJFAG3NFD47U",
      "comment": "Wallet16",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "9f9aNsmJxXgMZke5sRYFbfnH5fIFclSosqSl1mK4Vd8=",
        "vote": "h8ybeZLDhNG/53oJGAzZ2TFAXDXaslXMzNBOR3Pd+i4=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "EQ5XMOLC2JY5RNFXM725LRVKSTOHWBOQE344ZC6O2K4NW2S3G4XQIJNKAA",
      "comment": "Wallet17",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "R2LzBwBOEoMEcN6j2Pq9F1RKgrLrqnTyW/iT/tlIRZg=",
        "vote": "FnP52cIaWwqpJ6dE3KuM3WSGaz+TNlb/iM7EO0j7EZQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "3PUAOGK2PIEH6K5JTQ55SCV3E52KSLDPUAWDURMUNST6IIFCH347X5SNAI",
      "comment": "Wallet18",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "HfTcCIGCoAgUMCHalBv2dSC2L7XCPqPmCmWmxO26Vqo=",
        "vote": "knBY5MY9DkIguN41/ZoKvSGAg92/fhw64BLHUw0o1BU=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "DUQR2JOFHCTNRRI546OZDYLCVBIVRYOSWKNR7A43YKVH437QS3XGJWTQ6I",
      "comment": "Wallet19",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "DRSm3BAHOXLJLPHwrkKILG/cvHLXuDQYIceHgNPnQds=",
        "vote": "9G4AtYrLO26Jc3BsUfNl+0+3IjeHdOOSM+8ASj9x7Tg=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "NWBZBIROXZQEETCDKX6IZVVBV4EY637KCIX56LE5EHIQERCTSDYGXWG6PU",
      "comment": "Wallet2",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "0eG0edle+ejWcS4Q8DNlITgqaKqNvOtCxNQs+4AncGo=",
        "vote": "V4YUoGYXrgDjCluBBbBx2Kq9kkbCZudsuSwmSlCUnK0=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "U2573KTKRCC7I47FJUTW6DBEUN2VZQ63ZVYISQMIUEJTWDNOGSUTL67HBE",
      "comment": "Wallet20",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "cDT+xkHQJ13RgfkAUoNMfGk890z2C1V4HSmkxbm6gRk=",
        "vote": "r66g4ULatIt179X+2embK0RgwoLdPEq3R3uTTMfP9Hk=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "DBGTTXBPXGKL4TBBISC73RMB3NNZIZBSH2EICWZTQRA42QKNA4S2W4SP7U",
      "comment": "Wallet3",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "DmlAnKrkD8lgUB1ahLsy/FIjbZ0fypaowyDc8GKwWZA=",
        "vote": "ROBSmA9EfZitGyubHMTfmw8kSiohADB3n4McvTR8g88=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "XKZWM4PWPLZZWIANNT4S7LU26SPVIKMCDVQAAYRD4G3QJIOJL2X6RZOKK4",
      "comment": "Wallet4",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "74a0jcs/Y/uCh24vej1rb6CHu64yvW2nYrM0ZUVEhMo=",
        "vote": "rwkur9iwJbzNECWvELxzFeJpbZl7dpiThgPJOHnRykg=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "LPBKDDUNKPXE7GAICEDXGTNCAJNC6IFJUSD4IK2H2IIB3OAFXLM3RLLIVQ",
      "comment": "Wallet5",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "V4ldV+IY068YK/h7Wb6aNRIo8pr2bYQg8KDgFd25xVw=",
        "vote": "d2KdyajjKvpukuGmM2MxEC9XDEgjjF/Spsevjd877RI=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "MZZS43WEFY56LV3WXEVLROT3LYFLEBZ536UY3Z3J56S7EI3SYYOJVO6YRM",
      "comment": "Wallet6",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "BoBmrNpHTxySZ8DIlg5ZlINKwTPd/K75CCdhNzs9alo=",
        "vote": "N6v+PVEUn9fLZb+9sQDu5lpCpsXLHY0skx/8bWDqk7Q=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "RP7BOFGBCPNHWPRJEGPNNQRNC3WXJUUAVSBTHMGUXLF36IEHSBGJOHOYZ4",
      "comment": "Wallet7",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "n0LW+MxrO2S8/AmPClPaGdTDC5PM/MENdEwrm21KmgU=",
        "vote": "/e1z3LMbc8C4m9DZ6NCILpv7bZ/yVdmZUp/M32OSUN4=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "RDHKWTWXOE5AOWUWTROSR4WFLAHMUCRDZIA7OFBXXMMRBXGQ4BYQRPOXXU",
      "comment": "Wallet8",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "AGJ4v2nOA62A8rGm4H56VEo/6QdhVVJUuEASUybDPNI=",
        "vote": "eL2GxfrIoG2kuknlGa8I6vPtMbpygYflrye0u/hE4Lg=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "UXPVVSG7EYC7YR7PRVOZKWYYYZPKEXWGYR6XHBMSAV6BHKQEVFYVYJBVQI",
      "comment": "Wallet9",
      "state": {
        "algo": 320000000000000,
        "onl": 1,
        "sel": "P4tRdjhyJ9dSNItTY+r2+tQmPfHa6oBAzIh4X3df4gM=",
        "vote": "VHITXAytk0804xXBLBVKGlRAcAcDSZKcR2fiz4HtWBU=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A",
      "comment": "bank-testnet",
      "state": {
        "algo": 200000000000000,
        "onl": 1,
        "sel": "r6aMJIPeqUPB8u4IvOU/wihF+sgqJVsjibvsYHVqj1s=",
        "vote": "mPB1VDBFOPSIEFhXo7VJRLxn45ylDSRnO8J1nXQf4f0=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "GFICEF3GYRENRQHINLRPG7TS7TUIOARUIN7KWXWFROSG55BWFFRCRX5DAA",
      "comment": "n1-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "38qDzZjLPfernXNx7leElHsl39WLXMSgLHbEACeNgn4=",
        "vote": "8ITl30j5PTSDjmR26G3/rZL7IQM3cSfqqxnJSZf3X0w=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "GFY7ND6YSM5OGNSMAJDYCO6O75SWQRCYOJHCWOPYHUYCWQFWML52TWREBQ",
      "comment": "n10-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "iwwKBjoUUUePkoG0ldxc0v6i1fIhVySn2l2kWwekn2A=",
        "vote": "DaZFFz72XkcUIuPXcEz6VxWj4SVjzMpOwpTfO2k308g=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "VQFEAD2SXHMLJ3BNSGYUHRZZWBOI7HUQZGFFJEKYD3SGNS667FTMPRDC4Y",
      "comment": "n11-testnet",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "ckpVY6EaDInNeU1WLHQQXNsAaQnh+bpFhzNWzw0ZirI=",
        "vote": "4N1HJ9R2TrTEzLOyO1vUWPYi6sUcdAwQWoHQNBR/CME=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "6J7K7FIYKWTT3LSOZKYWAMSZC5RDID4CJ24C2S5DBQ5V7YUIHOBHPAO4KY",
      "comment": "n12-testnet",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "n16osH+x1UIrzDNa7PCZHn/UtheRoLcTBwGRnx0fTa8=",
        "vote": "Tj0inLse0V3sQRPw+5rVQTIWOqTxn7/URDzUaWGHftg=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "G5KCM3LSFV4GHRYQBXGWTNMR5XESE3PIRODD7ZLASPIGOHPV7CO7UKLZFM",
      "comment": "n13-testnet",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "tveXF/sDXqBXQY52IEMuvTeVguKzPfN8GLdKgtv3gRg=",
        "vote": "uwQJnVuqEtdGnWbbfu+TTLe++56z8wQCzv22IDioALE=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "XNQAMZMMLQV3TGYGJYYLYZUHP4YNEKAJM6RAMJ5SBXFLS3XDBIUVGCZPH4",
      "comment": "n14-testnet",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "8xotecjUoo1YVzWME3ib9uh+kPUNnzsFcuHrjxxhjZM=",
        "vote": "oQ/iakoP5B6gTTm0+xfHHGFS4Ink30I6FWUGkxRNfo8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "WXCLU5C6QH6KPVNAHNBGFUMC5JAOQCZP3HF76OT2TH3IAI3XTSPCLVILSU",
      "comment": "n15-testnet",
      "state": {
        "algo": 200000000000000,
        "onl": 1,
        "sel": "NRxs0rM5dov2oZrf6XrFSmG9CRlS3Bmzt0be7uF/nHw=",
        "vote": "R8xKtpYYNuTuTqMui/qzxYpc1m8KpbaK/eizYxVQDaY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "7NRVO2ABPGFRX3374TIJZ46BR72CCSHKTR6PG5VVYNLUPWUVXGOU3O5YQA",
      "comment": "n16-testnet",
      "state": {
        "algo": 200000000000000,
        "onl": 1,
        "sel": "IQG+jgm2daCxMLxm/f9tTVrDk/hD0ZhB5dxDQn47BSE=",
        "vote": "CGwAHrq3QFFlsP7NmHed+Xx4BwFsE2f6dB30Os75KxY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "537URFEXANB7M6UVND6WDM75DPRRORDXWFLSOG7EGILSKDIU4T32N4KAN4",
      "comment": "n17-testnet",
      "state": {
        "algo": 200000000000000,
        "onl": 1,
        "sel": "SdLlaWBe8B1JanMq0Y7T1Z9C8dKhI36MQiSffXQt7Lo=",
        "vote": "k4Xr6Bg6VpcY0GKwfr6kI89KqOihmCOToLLuIgFjv9c=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "ZNQXW7V5MISZFOZGVLAHKXS7GLWLXCLRPZTTIAZSTFRZPYTC54NWDZ6XZY",
      "comment": "n18-testnet",
      "state": {
        "algo": 200000000000000,
        "onl": 1,
        "sel": "TNMELlR1C+r4OmGVp9vc9XlehgD3a0EwfrepuMiDe+c=",
        "vote": "060veVAG/L2r2IAjqs2TcYy2cthocqrhgrCCoP5lzZ4=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "G3WQEPSGZOQVVJ2H3F6ICMHRIE2JL6U3X3JDABWJRN4HNDUJIAT4YTOGXA",
      "comment": "n19-testnet",
      "state": {
        "algo": 300000000000000,
        "onl": 1,
        "sel": "ktbtHTm1mUU5u/VMrOuMujMgemUf496zilQsGBynsxQ=",
        "vote": "XHXYdLvxKIIjtlmwHVqxvtAyRDE+SQR1tpzgXoNo5FA=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "2YNZ5XDUHYXL2COTVLZBRYV2A7VETFKQZQCPYMQRBOKTAANHP37DUH5BOI",
      "comment": "n2-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "u7lR9NcWfssuMvFYuqCi5/nX0Fj9qBKbE0B2OpRhmMg=",
        "vote": "/UGQ/1dcp7OTmguYALryqQYRj0oMWhs/ahAbQTL/mRA=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "IH5Z5UZCZKNAH5OICUGHFEYM2JDMJRUSIUV4TZEQYHRNS3T2ROOV32CDIA",
      "comment": "n20-testnet",
      "state": {
        "algo": 300000000000000,
        "onl": 1,
        "sel": "Jbcg+BVB6EOTe42U0dq1psQfoFZItb6Phst22z33j60=",
        "vote": "8Y1QY+WJIziffLecmnr0ZRGJFKtA3oVALQoD3nVKlt8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "FFJZOPQCYSRZISSJF33MBQJGGTIB2JFUEGBJIY6GXRWEU23ONC65GUZXHM",
      "comment": "n3-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "+K8AsLfvuTEuHMANNp2LxGuotgEjFtqOjuR/o4KR6LA=",
        "vote": "SerMKyY37A1jFkE0BdrP+vuTdVn9oOJc5QjC5f98Dz8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "ZWYIEI37V6HI62ZQCPJ5I6AIVKZP6JVCBQJKZEQQCWF4A4G2QGFENKS5XU",
      "comment": "n4-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "SmhBpQdh23++6xC01unged2JU1Wgm2zZ8v5LQiG/VqA=",
        "vote": "U2lZo9ahjkKBvcS3qSWsmSx+PGI/m6OtnQrQOH1iuII=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "V32YQ6LMMT7X6MML35KOX4MKY7LXWEH4JETZYKAXQ5RX4ZQQ6FAJJ6EGJQ",
      "comment": "n5-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "0yRtE7WSj32D5e/ov4o22ZgipQvqJZ6nx9NX1LdxFJI=",
        "vote": "scoN8x6Eq0bV4tBLT5R59jU+8gmHgh/6FX6mfV2tIKY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "OEFWPZHFT25CSDHFRFW62JANGQLB5WD25GJBCGYTTPHFUMAYYD7SEAIVDI",
      "comment": "n6-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "dWChUcA1ONX3iNEvHu9GST67XRePhAv6jd3XWt5clvI=",
        "vote": "rTfQ/l3lEfGQtzwjFii5ir2nCLSU+RT+0xI5af/XDEU=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "KCQLDL4GCVDLDYW5PYK7GJTUGHYRJ6CZ4QSRIZTXVRUIUAMDKYDFNUIFHU",
      "comment": "n7-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "gNXMo6XiZvuQs2mtomJZtra7XiZHySIOWLuWivu4iso=",
        "vote": "okgQcI/L7YDAMOyqrLKs6CUB91k+mMFfMTaEb+ixvyY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "UMMQNIYQXSI4VBGBXJUQ64ABURY6TPR7F4M5CMCOHYMB7GPVIZETZRNRBM",
      "comment": "n8-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "ukzMIkE2U33xKq6LGX19NBLirZNANQAf3oiZtlkn5ls=",
        "vote": "HYHBaeVeN0DXYBNjRBuGtZqrBr3bSBC1YDQrv93dNrc=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "2INEY2MWIWIUNQS24YVXKT4M3RIKMEZGTVAOJG47N7EOJE7MKXOC6GJSMU",
      "comment": "n9-testnet",
      "state": {
        "algo": 150000000000000,
        "onl": 1,
        "sel": "7aUtPCawOYpPYjVd6oZOnZ+1CZXApr8QR4q1cOkVyWo=",
        "vote": "kcq1XWHnMrjbv/fvMmzIfGZzDtJtdL7i70lpWZ0kGi0=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "IE4C3BNWT4EYKPUZXGWDOOKBTJFVOYAZKBCWFYRC37U7BJKBIUH6NEB7SQ",
      "comment": "pp1-testnet",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "C3PdYqoDjrjyaGvZ6M/W0E56Mv5BXdtRwj7+4unpxDM=",
        "vote": "8fdNikU3nMNyZb3AZlNTnsfsytvrd8bK2b/dYQgJj30=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "7WCI7XPEMWY6XNWHG2VXGYGDLHPTJ333CZ2WBGGUHCSYPTXPBWYCHZYTSE",
      "comment": "pp2-testnet",
      "state": {
        "algo": 25000000000000,
        "onl": 1,
        "sel": "l3K4aA15T42mTM+QE7GpOzbOcth6hMljBxna7gSR8IA=",
        "vote": "NsjSVQJj4XxK5Tt0R7pvU6wQB0MRKHDwC9F2bfUX/vM=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "WYX5JGDYM7TBTMBBEE2OI4GC4KVCTLB2P67B3PUQQS4OMUERE7NIIZDWO4",
      "comment": "pp3-testnet",
      "state": {
        "algo": 25000000000000,
        "onl": 1,
        "sel": "YmLs97jSdlbYU1H0PwZdzo6hlp0eyBwJ+ydM9ggEENI=",
        "vote": "GeDnbm9KKEu2dZ1FACwI0NsVWgoU0udpZef06IiTdfQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "2GJF4FEEPNCFKNYSOP6EOQGDQQCGDXPQHWE474DCKP5QO3HFBO73IBLBBY",
      "comment": "u1-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "NHZ3VOL34MVWENM72QB6ZBRDMFJTU6R57HAJALSBERH4BNAGR4QDYYBT7A",
      "comment": "u10-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "PTLGEQAIGTDWHPKA3IC5BL5UQE52XDZHQH7FUXRV4S6ZBRR5HGZENQ7LTQ",
      "comment": "u100-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "3IE2GDYYSI56U53AQ6UUWRGAIGG5D4RHWLMCXJOPWQJA2ABF2X2OLFXGJE",
      "comment": "u11-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "IAMUOCM2SEISQZYZZYTLHKSAALDJIXS2IQRU2GPZUOZWB2NLMFZPJSQ7VQ",
      "comment": "u12-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "54GKXNGS7HNFHZGO7OIWK3H2KPKZYWSARW7PV4ITVTNCA65K6ESRKI6N3U",
      "comment": "u13-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "5ZSFGF66FIJMMRORTYD2PLDAN67FA2J7LF3IYF4ZKD4DJHLEBYJ76DXGVU",
      "comment": "u14-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "DY7K3FLRZTW2ZTYVOC4TCGK4JBL7NSJ4GR4BU252QNAVOCVTGEBCPCSJME",
      "comment": "u15-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "JG4JQZNYP2524UDVRPPIMSFCIVQPVXLB5AKHM76VXIIRFNMIN3ROIYW65E",
      "comment": "u16-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "7J4QX5DVIXSWBC2NJB44LPPUJXOAJQFMBCOS4EDI3XOE5WS76IY7WFTBQI",
      "comment": "u17-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "6SA2WG5XM5Q6SSMBRK3TOHY552A75RVANBQQMKTT67PLUN44T3CJZAQOPM",
      "comment": "u18-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "64DCC5CMTM4SMMO3QRTY3EDCHS73KDSNNH2XZL262DBK2LR4GJRETWUWIE",
      "comment": "u19-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "TQ2B4MTCC6TARNEP4QPPMCKNBBNXKFTQKPVLAFC5XXRR2SWV5DICZELJOY",
      "comment": "u2-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "ATNCIRLQLVZ7I4QBGW54DI6CY4AJVBQBPECVNS645RBMYDTK6VV55HXFUU",
      "comment": "u20-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4LP77VEVJ7QNESED4GICPRBZUNP7ZLKKLEVBRDSKX5NZSUFXPSEA575K5E",
      "comment": "u21-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "7D34RBEHVI3A7YTQWOUTCSKNQYS5BDBN4E647DOC6WDVOLHPDPSSBY4MWI",
      "comment": "u22-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "UMMKTTPNHIURGX24K7UYJ7T3WBB5J7OYBOQJ5WLPRG3BDYWJAEJLVBNHME",
      "comment": "u23-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "EOPSQC3QTL7QJ4AQ2J4OJIJMKQLTMIEETJI7OFWYADIMHDWMHQ6MWCTUMQ",
      "comment": "u24-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "XT3AVLURALOWTIMGZKB37J2M22NUQCRXTL4DJZHSTPCGLNQKVL7MR3MKFM",
      "comment": "u25-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "WS63FDTLLYHC2NS7NXTEO7RPLNMAFM2D2BPJLTMAQJWPR2JCNYTTRMSOAE",
      "comment": "u26-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "P5S5GGUHOMVOKWOZPJO74MBYVRXQWDBW6AOTHQZVKJKFGM7VBU6CNR4ATI",
      "comment": "u27-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "PXVAI3MUYH4WWJXEQP7XNH3YIMO5ZBAFJWYUL7DOGPAHALE4K6GZBF4THU",
      "comment": "u28-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "VGTKWLFANSULZAFDGBONHF55VVKE4V4F63JRDB66XM4K6KCQX6CL22WPRE",
      "comment": "u29-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "QB2OTQ6DKUEJFP66A37ASIT4O3UZUOX24DAMWU2D3GCBDIYIXSIDHSXO4E",
      "comment": "u3-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4F6LA64ZLFN33ATWJ74UPAX56OLTXPL74SS5ATXUL7RGX7NKEFKMAWUQYE",
      "comment": "u30-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "3JBNL7BZECXKYWZRPWETNL65XEYMAHLC6G3MZN2YMPFL3V7XSDXZEMBHVQ",
      "comment": "u31-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4M2QSKTXKPPZMNUAQ4UDS7ASMQCEUE4WTWGV6AM326425IJ64UNZBCIRGA",
      "comment": "u32-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "J37V3LXHPRRKBODXNMNYNUJQIICCFFC4O4XB4YJCPVUAVZNOUG5DWDCEIA",
      "comment": "u33-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "I75JBQHNYEYM3J742RBVW4W6RR3YY3BLG2PKO4PXYLVNEX5L646ASDJOOY",
      "comment": "u34-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "ZHEIOZ7E2BEBCCKK5QM7DCZAOPTTONMQWHNJ6FOLKBHY466VON6DCZERD4",
      "comment": "u35-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4QMGP4C6OMSCNJI25H7UQGBFHRHL7KXAEQI57JNAXEO2EW3VT6D6LODT5Y",
      "comment": "u36-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "KRED3JOLOJE3SLL5NGHAWSUGEMHCYJLD6PX43SIJYN2GC6MS6HPUPPO2LY",
      "comment": "u37-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "SVFLDISKS4PDMJKOB6DVVVN6NQ776FHZMGWCOUQVQCH6GXTKCXIHTLYRRQ",
      "comment": "u38-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "7IWGAPZ4VWRZLP2IHFSAC3JYOKNAZP6ONBNGGWUWHAUT7F23YFT3XKGNVU",
      "comment": "u39-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "O2QVZMKATOIEU2OD4X42MLXAYVRXLRDKJTDXKBFCN3PCKN2Z3PUS5HKIVA",
      "comment": "u4-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "H2YN73YPRWKY4GT744RRD65CXSQZO7MK72MV4RDHTIBV6YQUB2G56TVF2Y",
      "comment": "u40-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "WGUAFWHRRX7VXPO3XXYCJL5ELO6REUGD57HRMBKTALT2TTXOLSHNOUEQCE",
      "comment": "u41-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "RYHCD7GPAEBRV657FJJAG2ZZUDVPR66IU7CA5Y7UDMYSEEIWR4QDNSPLYQ",
      "comment": "u42-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "BKTO5TKB4L57YWTZKQBOQ37EWH2HVXGJPXP3L6YSYOAWP3CYYBWLZ2PHTQ",
      "comment": "u43-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "FL7LZ57VQQNW5NDJK2IKEAHIXRTB7VFBJEA2MIAEK3QVZPIBGLYW7XSZDY",
      "comment": "u44-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "MXXQXZS2TAMIULLXXLX6MM6AHJAOQLHEIB2U3LR4KYKK7ZKRVUSHTU62QA",
      "comment": "u45-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "UGOPPKTJQ2KPHU5I56733IMT3B7ECT5O44GW2FYX5SNDVIEDG72Z5GC5IA",
      "comment": "u46-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "Y7MGWPRBHQN2PF3I2A3RWCQMVA42VR6FJONJ3W26WGKE4KMCGCVJIDLHEY",
      "comment": "u47-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "NNFIWU43AUEZIUIQQECDXM3HRPUEJMPPZLXTM4ZFJKHWSZ2FEGCVMMJUBQ",
      "comment": "u48-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "RN3HTSJKSUO6OECM3OPDFQQ2FYZWEY2OWAQGSMQSGY4DI7JJ4HBV2OIJJU",
      "comment": "u49-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "OLYQUMZKLYDX2FVHECURBX4SRQSLMIIWN7D7VRJG7B6DS3IU6M5WYVNAAY",
      "comment": "u5-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "PIG4P6JA2WDG7HBBR4FFDMVUCUD5Y5CTQ3K3KY34Y4AMT3CWEMVIKQLZZI",
      "comment": "u50-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "QIDX47JRS37LRIYVY744SV7KTFGYXY5ABEK2VALNZCMN2H4FBLO7WWKYRM",
      "comment": "u51-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "VQZCPUMOYIGCXOK2AK4XYYLWJNRBLS457IL4OSBKGVBHFZ5QPLTCUOTW4A",
      "comment": "u52-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "WE2AIYHXI2LHABITCPTZRBTLFT54HPL4MKIR4HTASARNGCCZLXXDE67H3M",
      "comment": "u53-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "HAIGM3LXXVKDCGCNQELNOBFZKP6C4A2ZY464F4TB7GWSVDN6I4SI7EOZUE",
      "comment": "u54-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "MVZLGPXT6DZQIORE4PIO7NZD7QMJOZZZCOEVPZ3EQX2V4WG3PFU3BXUGDI",
      "comment": "u55-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "MB5XJGVVKQU7NSEWWP65QW6H4JVEQYPA5626J4NGQP2E4BUMXRTEGW5X5Y",
      "comment": "u56-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "EODZLNWFSRYZKLLF2YAOST2CYQCBRQGXPFQJLDW4CCMYFTYKBSWMF6QUAU",
      "comment": "u57-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "LPAMNP7GJC5CNOMWRDII47WWYPF3TOVEIBDSSJA6PKOCPZ5AKRUWMIU2OM",
      "comment": "u58-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "THRYS4MAIMEKG7BSAZ4EOKCVUJ7HA6AOCTK2UOKDGZ4TF7Q4BRVTBOUSYU",
      "comment": "u59-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "7V7YITMPBTJ3IHHS2D35PVWRZGNFYWWQVRMTI4QP2CBPSKNDRGG66W2HFQ",
      "comment": "u6-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "FHA2V46TK5CW66HQPOMLTH5PSKX2JX2IWLWZIYJUZ2RI7SK6HSSBTJBNHM",
      "comment": "u60-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "7EJAXCVH7XLWDCWSXID4FNZ6T2SZRA4S7XIZOWA74ITAB272ZF2T5LSWSE",
      "comment": "u61-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "K5L3YNZPU6SVNJOWAOKULCWBPIBNMR2VBCASVI4NWDM2APZ6GL36DFDR5Y",
      "comment": "u62-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "5BY6RFBNUYHBYH4E4AWVMEOMI7YFKX7X3IPB5GRGAHH4BSXHIL34P3H43A",
      "comment": "u63-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "BX2UBG5VCT2ASTGXHVG5NS6VVCYVB6GLKBN4NAAN7ABSTP7BMYCX2T2WEY",
      "comment": "u64-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "37JPBYKXMWF6DO3FFWW53LBQCG636MTC7WG6DTRAPDFVXUIATFOMFR5ZLQ",
      "comment": "u65-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "ODSPT3NISYMGEE3TJ6U6JCVC44L7DUCPHIV2QMPPRKBWJDALALGVCAPMRE",
      "comment": "u66-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "CQA775B5TCU72Y2BNL6VCURBVJE45QV77RXHQ5KYRMMP6NCQ5BR7XJRYRA",
      "comment": "u67-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "3Q4SYOBDOAVXUUTKBXEFFSK3BQMUQX5ORZPDA4PHB56KJJONPFFJ7YZ6HU",
      "comment": "u68-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "K23ME4QVDHSJWMGUHPGCL2OODAGBHIBW2KGYLLIR3UAEFD5ZW2KFB4WJ34",
      "comment": "u69-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "UD2OLL24RFDFMAKK7CCHKFIABPAP7ET4CYQUEYCJVGEIEJUAMDOGJZT26Y",
      "comment": "u7-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "KYXWZODLYDHTDMRUBOGOEV42G6H6KJ2JSBFZBP6XNWT42A6QEMEW23JWAM",
      "comment": "u70-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "V464X6D3XJVVJ372FFC2NBBDZLBNQA6H55J57WJMMSNOLHOJQ5UF3EUGNY",
      "comment": "u71-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "K27ODRPQARZM3236D2XC27QIV27GO2MUR65RGAJKO7UACIFYHG5QKPOCFU",
      "comment": "u72-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "FT3JD6TXUZOLOMN4O5CFZYSIHR4T5XJIF2YNV6WGEORNO2X65QW3VUP77I",
      "comment": "u73-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "WOTGZ4WOQ4S7YWVAOQ52GGOQPYQI2M7EPZENR27AOZLYFIEJDI3RYFB7OU",
      "comment": "u74-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "46MGTGNCTAC62NVNAVXAGP7PUJJIW5GXYYTSUDURCBSRZEDLGME7ICGE4E",
      "comment": "u75-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "MBTXWM5M5XQNUEKLBTW7GPU4LFPUETQQPVUBRCOA7FQ47H4J727NFRKKQE",
      "comment": "u76-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4MCTFKPQCY25X6QARHGVD75OYUMQAAU5QLWCE2EM37NWOS7IFJSABMGKBI",
      "comment": "u77-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "PY6K3OGCXZVYQYZVA7W3MVZCAU5AFAWQ5J5THILXYIBYCKCGH4ELFU6TNU",
      "comment": "u78-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4ABEMED4I7UYU6CJSLWYQXQHOK2XCQ443BSHR3SL7QJGXNYJ5QCYILSSNU",
      "comment": "u79-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "AXBINC5KA3I6IF3JAMKYQU3JLYTA5P2U4PUW3M4L53NEBNCRLHDHHOT2HY",
      "comment": "u8-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "5SXA2C6CGZ63OYDY5G4NFLIPJLKCZAMQWLMD2CBNSHUEXVS3ZYHAQCI5TI",
      "comment": "u80-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "EVP6MJIZWN6EE64TKEI4ANETP25MHYVXFWESU626TFA5VDVC75KSBGAA54",
      "comment": "u81-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "QAUV22GPBAOCO2JGAJF7U474S5SKXVWSZ7KG6P22P4MH3GNBGEJXAVDQLM",
      "comment": "u82-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4FOOFGIWV4H7AXTEJXV2C4ONZ5NXAMUDKJSZDLSKACZ4JA4SWIU6UTLZAU",
      "comment": "u83-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "ARUMRBUW3HBQXE4QAL25PPVWAJSKGORTNUIOW3VA5GAMDECOVNYC7GJJS4",
      "comment": "u84-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "EJGCYTFUZPJDL2JBZJFQXKZIYJUDB7IBF3E2BH6GXWYWXUHSBCKYFJUKSU",
      "comment": "u85-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "FJMEFROCSGQ7C7IXMAPUST37QTQ2Y4A7RMLGK6YTUGHOCLOEL5BDE4AM2M",
      "comment": "u86-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "4V635E2WOGIKKWZ6QMYXDWQLYTUKRN7YAYADBQPETS75MKCR66ZC5IEG5M",
      "comment": "u87-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "AFJB4HWJLTMMA45VZAJJSUOFF7NROAEEMGT4Z3FQI5APWY472SJ6RNBWU4",
      "comment": "u88-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "BYO56YQOSRBUTNPXYO4XDMG7FU7SIP3QGVKAYQIJVJ4UIIMBRG3E4JMVD4",
      "comment": "u89-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "FQJO4LDTXEVQ2ZBFYDEAOYPQQZCZTMASMSXJ6V7LBYKOTFSCBUKKIU3DXA",
      "comment": "u9-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "WUCEVFNJGUNLMNG2AJMVYJRGQUFXRAFVX2ZRT7AC47WS6IRHPXHSUZ4NUA",
      "comment": "u90-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "5J5Q72IHCVAK5NE54ZI2RUZUF3HN2EAQEYQ674H3VX4UUHBMRYAZFRQDIY",
      "comment": "u91-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "2LK2SZ3L4PWUXXM4XYFFSCFIV7V5VQJUDFVK7QXK6HJL4OUQKQLWG77EUI",
      "comment": "u92-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "QUWHMJLFQAIIG5LV7NK5VNESUUW23RINBSHKKKQDIV4AP56RSTYSNZHDRQ",
      "comment": "u93-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "YJEGUEJ2UW2ABLO6XI5QIHQID5ZKUDUDQPHQEN7MH5SS2FLZ573CHRHCZM",
      "comment": "u94-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "XOUVBGEZMDVYPES4MGTAEBYU5O6LOCOH27ZJ3ML7ATWEU63N6IWW6F4BLM",
      "comment": "u95-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "6CFS2YVK2IMVVFBGGHSPUQBIKMNWRRB44EIUUB4EFXAL7IOJXAHRGXKAGA",
      "comment": "u96-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "W5ITKFRKK265A4WKF7IRCZ4MCC7HM3INCJGKPPH3AEKDFYMOJJ4FDLQWYI",
      "comment": "u97-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "MBMU3IODI6OFX34MBDMNTD6WSVA6B3XLDVB3IHZJQY3TZUYBPKRNFTUQSM",
      "comment": "u98-testnet",
      "state": {
        "algo": 2000000000000
      }
    },
    {
      "addr": "CKNVTB7DPRZO3MB64RQFPZIHCHCC4GBSTAAJKVQ2SLYNKVYPK4EJFBCQKM",
      "comment": "u99-testnet",
      "state": {
        "algo": 2000000000000
      }
    }
  ],
  "fees": "A7NMWS3NT3IUDMLVO26ULGXGIIOUQ3ND2TXSER6EBGRZNOBOUIQXHIBGDE",
  "id": "v1.0",
  "network": "testnet",
  "proto": "https://github.com/algorand/spec/tree/a26ed78ed8f834e2b9ccb6eb7d3ee9f629a6e622",
  "rwd": "7777777777777777777777777777777777777777777777777774MSJUVU",
  "timestamp": 1560210455
}`

func GetDefaultTestNetGenesis(ctx context.Context) model.Genesis {
	var g model.Genesis
	if err := json.Unmarshal([]byte(DefaultTestNetGenesis), &g); err != nil {
		runtime.LogErrorf(ctx, "error could not unmarshal default testnet genesis json: %s\n", err)
		return g
	}
	return g
}

func LoadTestNetGenesis(ctx context.Context, path string) model.Genesis {
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
