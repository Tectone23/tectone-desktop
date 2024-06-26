package genesis

import (
	"context"
	"encoding/json"
	"os"
	"tectone/tectone-desktop/internal/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const DefaultMainNetGenesis = `{
  "alloc": [
    {
      "addr": "737777777777777777777777777777777777777777777777777UFEJ2CI",
      "comment": "RewardsPool",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "Y76M3MSY6DKBRHBL7C3NNDXGS5IIMQVQVUAB6MP4XEMMGVF2QWNPL226CA",
      "comment": "FeeSink",
      "state": {
        "algo": 1000000,
        "onl": 2
      }
    },
    {
      "addr": "ALGORANDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIN5DNAU",
      "comment": "A BIT DOES E NOT BUT E STARTS EVERYTHING LIFE A MANY FORTUNE R BUILD SIMPLER BE THE STARTS PERSEVERES FAVORS A ENOUGH RIPROVANDO POSSIBLE JOURNEY VICTORIA HE BOLD U WITHOUT MEN A K OF BORDERS WHO HE E RACES TOMORROW BUT WHO SINGLE PURPOSE GEOGRAPHICAL PROVANDO A KNOW SUFFOCATES NOT SCIENCE STEP MATHEMATICS OF OR A BRIDGES WALLS TECHNOLOGY TODAY AND WITH AS ET MILES OF THOUSAND VITA SIMPLE TOO MUST AS NOT MADE NOT",
      "state": {
        "algo": 1000000,
        "onl": 2
      }
    },
    {
      "addr": "XQJEJECPWUOXSKMIC5TCSARPVGHQJIIOKHO7WTKEPPLJMKG3D7VWWID66E",
      "comment": "AlgorandCommunityAnnouncement",
      "state": {
        "algo": 10000000,
        "onl": 2
      }
    },
    {
      "addr": "VCINCVUX2DBKQ6WP63NOGPEAQAYGHGSGQX7TSH4M5LI5NBPVAGIHJPMIPM",
      "comment": "AuctionsMaster",
      "state": {
        "algo": 1000000000,
        "onl": 2
      }
    },
    {
      "addr": "OGP6KK5KCMHT4GOEQXJ4LLNJ7D6P6IH7MV5WZ5EX4ZWACHP75ID5PPEE5E",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "AYBHAG2DAIOG26QEV35HKUBGWPMPOCCQ44MQEY32UOW3EXEMSZEIS37M2U",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "2XKK2L6HOBCYHGIGBS3N365FJKHS733QOX42HIYLSBARUIJHMGQZYAQDRY",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "ZBSPQQG7O5TR5MHPG3D5RS2TIFFD5NMOPR77VUKURMN6HV2BSN224ZHKGU",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "7NQED6NJ4NZU7B5HGGFU2ZEC2UZQYU2SA5S4QOE2EXBVAR4CNAHIXV2XYY",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "RX2ZKVJ43GNYDJNIOB6TIX26U7UEQFUQY46OMHX6CXLMMBHENJIH4YVLUQ",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "RHSKYCCZYYQ2BL6Z63626YUETJMLFGVVV47ED5D55EKIK4YFJ5DQT5CV4A",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "RJS6FDZ46ZZJIONLMMCKDJHYSJNHHAXNABMAVSGH23ULJSEAHZC6AQ6ALE",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "AZ2KKAHF2PJMEEUVN4E2ILMNJCSZLJJYVLBIA7HOY3BQ7AENOVVTXMGN3I",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "CGUKRKXNMEEOK7SJKOGXLRWEZESF24ELG3TAW6LUF43XRT2LX4OVQLU4BQ",
      "comment": "",
      "state": {
        "algo": 300000000000000,
        "onl": 2
      }
    },
    {
      "addr": "VVW6BVYHBG7MZQXKAR3OSPUZVAZ66JMQHOBMIBJG6YSPR7SLMNAPA7UWGY",
      "comment": "",
      "state": {
        "algo": 250000000000000,
        "onl": 2
      }
    },
    {
      "addr": "N5BGWISAJSYT7MVW2BDTTEHOXFQF4QQH4VKSMKJEOA4PHPYND43D6WWTIU",
      "comment": "",
      "state": {
        "algo": 1740000000000000,
        "onl": 2
      }
    },
    {
      "addr": "MKT3JAP2CEI5C4IX73U7QKRUF6JR7KPKE2YD6BLURFVPW6N7CYXVBSJPEQ",
      "comment": "",
      "state": {
        "algo": 158000000000000,
        "onl": 2
      }
    },
    {
      "addr": "GVCPSWDNSL54426YL76DZFVIZI5OIDC7WEYSJLBFFEQYPXM7LTGSDGC4SA",
      "comment": "",
      "state": {
        "algo": 49998988000000,
        "onl": 1,
        "sel": "lZ9z6g0oSlis/8ZlEyOMiGfX0XDUcObfpJEg5KjU0OA=",
        "vote": "Kk+5CcpHWIXSMO9GiAvnfe+eNSeRtpDb2telHb6I1EE=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "M7XKTBQXVQARLS7IVS6NVDHNLJFIAXR2CGGZTUDEKRIHRVLWL5TJFJOL5U",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "Z5gE/m2E/WSuaS5E8aYzO2DugTdSWQdc5W5BroCJdms=",
        "vote": "QHHw03LnZQhKvjjIxVj3+qwgohOij2j3TBDMy7V9JMk=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "QFYWTHPNZBKKZ4XG2OWVNEX6ETBISD2VJZTCMODIZKT3QHQ4TIRJVEDVV4",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "NthIIUyiiRVnU/W13ajFFV4EhTvT5EZR/9N6ZRD/Z7U=",
        "vote": "3KtiTLYvHJqa+qkGFj2RcZC77bz9yUYKxBZt8B24Z+c=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "DPOZQ6HRYLNNWVQL3I4XV4LMK5UZVROKGJBRIYIRNZUBMVHCU4DZWDBHYE",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "PBZ/agWgmwMdmWgt/W0NvdTN/XSTrVhPvRSMjmP5j90=",
        "vote": "FDONnMcq1acmIBjJr3vz4kx4Q8ZRZ8oIH8xXRV5c4L8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "42GALMKS3HMDB24ZPOR237WQ5QDHL5NIRC3KIA4PCKENJZAD5RP5QPBFO4",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "p7axjoy3Wn/clD7IKoTK2Zahc5ZU+Qkt2POVHKugQU4=",
        "vote": "PItHHw+b01XplxRBFmZniqmdm+RyJFYd0fDz+OP4D6o=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "OXWIMTRZA5TVPABJF534EBBERJG367OLAB6VFN4RAW5P6CQEMXEX7VVDV4",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "RSOWYRM6/LD7MYxlZGvvF+WFGmBZg7UUutdkaWql0Xo=",
        "vote": "sYSYFRL7AMJ61egushOYD5ABh9p06C4ZRV/OUSx7o3g=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "AICDUO6E46YBJRLM4DFJPVRVZGOFTRNPF7UPQXWEPPYRPVGIMQMLY5HLFM",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "0vxjPZqEreAhUt9PHJU2Eerb7gBhMU+PgyEXYLmbifg=",
        "vote": "fuc0z/tpiZXBWARCJa4jPdmDvSmun4ShQLFiAxQkOFI=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "DYATVHCICZA7VVOWZN6OLFFSKUAZ64TZ7WZWCJQBFWL3JL4VBBV6R7Z6IE",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "KO2035CRpp1XmVPOTOF6ICWCw/0I6FgelKxdwPq+gMY=",
        "vote": "rlcoayAuud0suR3bvvI0+psi/NzxvAJUFlp+I4ntzkM=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "6XJH2PJMAXWS4RGE6NBYIS3OZFOPU3LOHYC6MADBFUAALSWNFHMPJUWVSE",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "PgW1dncjs9chAVM89SB0FD4lIrygxrf+uqsAeZw8Qts=",
        "vote": "pA4NJqjTAtHGGvZWET9kliq24Go5kEW8w7f1BGAWmKY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "EYOZMULFFZZ5QDDMWQ64HKIMUPPNEL3WJMNGAFD43L52ZXTPESBEVJPEZU",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 1,
        "sel": "sfebD2noAbrn1vblMmeCIeGB3BxLGKQDTG4sKSNibFs=",
        "vote": "Cuz3REj26J+JhOpf91u6PO6MV5ov5b1K/ii1U1uPD/g=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "I3345FUQQ2GRBHFZQPLYQQX5HJMMRZMABCHRLWV6RCJYC6OO4MOLEUBEGU",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "MkH9KsdwiFgYtFFWFu48CeejEop1vsyGFG4/kqPIOFg=",
        "vote": "RcntidhQqXQIvYjLFtc6HuL335rMnNX92roa2LcC+qQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "6LQH42A4QJ3Y27FGKJWERY3MD65SXM4QQCJJR2HRJYNB427IQ73YBI3YFY",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "nF3mu9Bu0Ad5MIrT31NgTxxrsZOXc4u1+WCvaPQTYEQ=",
        "vote": "NaqWR/7FzOq/MiHb3adO6+J+kvnQKat8NSqEmoEkVfE=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "3V2MC7WJGAFU2EHWBHEETIMJVFJNAT4KKWVPOMJFJIM6ZPWEJRJ4POTXGI",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "3i4K8zdmnf1kxwgcNmI3x50iIwAxDmLMvoQEhjzhado=",
        "vote": "wfJWa0kby76rqX2yvCD/aCfJdNt+qItylDPQiuAWFkQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "FTXSKED23VEXNW442T2JKNPPNUC2WKFNRWBVQTFMT7HYX365IVLZXYILAI",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "icuL7ehcGonAcJ02Zy4MIHqcT+Sp1R1UURNCYJQHmo4=",
        "vote": "tmFcj3v7X5DDxKI1IDbGdhXh3a5f0Ab1ftltM7TgIDE=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "IAOW7PXLCDGLKMIQF26IXFF4THSQMU662MUU6W5KPOXHIVKHYFLYRWOUT4",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "zTn9rl/8Y2gokMdFyFP/pKg4eP02arkxlrBZIS94vPI=",
        "vote": "a0pX68GgY7u8bd2Z3311+Mtc6yDnESZmi9k8zJ0oHzY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "4NRNE5RIGC2UGOMGMDR6L5YMQUV3Q76TPOR7TDU3WEMJLMC6BSBEKPJ2SY",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "orSV2VHPY8m5ckEHGwK0r+SM9jq4BujAICXegAUAecI=",
        "vote": "NJ9tisH+7+S29m/uMymFTD8X02/PKU0JUX1ghnLCzkw=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "E2EIMPLDISONNZLXONGMC33VBYOIBC2R7LVOS4SYIEZYJQK6PYSAPQL7LQ",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "XM2iW9wg9G5TyOfVu9kTS80LDIqcEPkJsgxaZll3SWA=",
        "vote": "p/opFfDOsIomj5j7pAYU+G/CNUIwvD2XdEer6dhGquQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "APDO5T76FB57LNURPHTLAGLQOHUQZXYHH2ZKR4DPQRKK76FB4IAOBVBXHQ",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "5k2vclbUQBE6zBl45F3kGSv1PYhE2k9wZjxyxoPlnwA=",
        "vote": "3dcLRSckm3wd9KB0FBRxub3meIgT6lMZnv5F08GJgEo=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "3KJTYHNHK37G2JDZJPV55IHBADU22TX2FPJZJH43MY64IFWKVNMP2F4JZE",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "o5e9VLqMdmJas5wRovfYFHgQ+Z6sQoATf3a6j0HeIXU=",
        "vote": "rG7J8pPAW+Xtu5pqMIJOG9Hxdlyewtf9zPHEKR2Q6OE=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "IVKDCE6MS44YVGMQQFVXCDABW2HKULKIXMLDS2AEOIA6P2OGMVHVJ64MZI",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "XgUrwumD7oin/rG3NKwywBSsTETg/aWg9MjCDG61Ybg=",
        "vote": "sBPEGGrEqcQMdT+iq2ududNxCa/1HcluvsosO1SkE/k=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "2WDM5XFF7ONWFANPE5PBMPJLVWOEN2BBRLSKJ37PQYW5WWIHEFT3FV6N5Y",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "Lze5dARJdb1+Gg6ui8ySIi+LAOM3P9dKiHKB9HpMM6A=",
        "vote": "ys4FsqUNQiv+N0RFtr0Hh9OnzVcxXS6cRVD/XrLgW84=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "EOZWAIPQEI23ATBWQ5J57FUMRMXADS764XLMBTSOLVKPMK5MK5DBIS3PCY",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "jtmLcJhaAknJtA1cS5JPZil4SQ5SKh8P0w1fUw3X0CE=",
        "vote": "pyEtTxJAas/j+zi/N13b/3LB4UoCar1gfcTESl0SI2I=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "REMF542E5ZFKS7SGSNHTYB255AUITEKHLAATWVPK3CY7TAFPT6GNNCHH6M",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "8ggWPvRpSkyrjxoh1SVS9PiSjff2azWtH0HFadwI9Ck=",
        "vote": "Ej/dSkWbzRf09RAuWZfC4luRPNuqkLFCSGYXDcOtwic=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "T4UBSEAKK7JHT7RNLXVHDRW72KKFJITITR54J464CAGE5FGAZFI3SQH3TI",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "eIB8MKaG2lyJyM9spk+b/Ap/bkbo9bHfvF9f8T51OQk=",
        "vote": "7xuBsE5mJaaRAdm5wnINVwm4SgPqKwJTAS1QBQV3sEc=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "YUDNQMOHAXC4B3BAMRMMQNFDFZ7GYO2HUTBIMNIP7YQ4BL57HZ5VM3AFYU",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "CSTCDvvtsJB0VYUcl3oRXyiJfhm3CtqvRIuFYZ69Z68=",
        "vote": "uBK1TH4xKdWfv5nnnHkvYssI0tyhWRFZRLHgVt9TE1k=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "4SZTEUQIURTRT37FCI3TRMHSYT5IKLUPXUI7GWC5DZFXN2DGTATFJY5ABY",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "THGOlrqElX13xMqeLUPy6kooTbXjiyrUoZfVccnHrfI=",
        "vote": "k4hde2Q3Zl++sQobo01U8heZd/X0GIX1nyqM8aI/hCY=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "UEDD34QFEMWRGYCBLKZIEHPKSTNBFSRMFBHRJPY3O2JPGKHQCXH4IY6XRI",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "jE+AUFvtp2NJsfNeUZeXdWt0X6I58YOgY+z/HB17GDs=",
        "vote": "lmnYTjg1FhRNAR9TwVmOahVr5Z+7H1GO6McmvOZZRTQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "HHZQOGQKMQDLBEL3HXMDX7AGTNORYVZ4JFDWVSL5QLWMD3EXOIAHDI5L7M",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "Hajdvzem2rR2GjLmCG+98clHZFY5Etlp0n+x/gQTGj0=",
        "vote": "2+Ie4MDWC6o/SfFSqev1A7UAkzvKRESI42b4NKS6Iw8=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "XRTBXPKH3DXDJ5OLQSYXOGX3DJ3U5NR6Y3LIVIWMK7TY33YW4I2NJZOTVE",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "5qe7rVoQfGdIUuDbhP2ABWivCoCstKbUsjdmYY76akA=",
        "vote": "3J3O9DyJMWKvACubUK9QvmCiArtZR7yFHWG7k7+apdQ=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "JJFGCPCZPYRLOUYBZVC4F7GRPZ5CLB6BMTVRGNDP7GRGXL6GG4JEN7DL54",
      "comment": "",
      "state": {
        "algo": 24000000000000,
        "onl": 1,
        "sel": "YoRFAcTiOgJcLudNScYstbaKJ8anrrHwQMZAffWMqYE=",
        "vote": "VQFKlDdxRqqqPUQ/mVoF8xZS9BGxUtTnPUjYyKnOVRA=",
        "voteKD": 10000,
        "voteLst": 3000000
      }
    },
    {
      "addr": "4VNSA2GZVUD5ZNO62OVVNP4NEL2LIEE5N3MZEK4BKH62KGKRLVINFZYTZM",
      "comment": "",
      "state": {
        "algo": 100000000000000,
        "onl": 2
      }
    },
    {
      "addr": "IVCEEIH2Q32DZNRTS5XFVEFFAQGERNZHHQT6S4UPY7ORJMHIQDSTX7YM4E",
      "comment": "",
      "state": {
        "algo": 408400000000000,
        "onl": 2
      }
    },
    {
      "addr": "PLFHBIRGM3ZWGAMCXTREX2N537TWOMFIQXHFO2ZGQOEPZU473SYBVGVA5M",
      "comment": "",
      "state": {
        "algo": 1011600000000000,
        "onl": 2
      }
    },
    {
      "addr": "KF7X4ZABZUQU7IFMHSKLDKWCS4F3GZLOLJRDAK5KMEMDAGU32CX36CJQ5M",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "BTEESEYQMFLWZKULSKLNDELYJTOOQK6ZT4FBCW3TOZQ55NZYLOO6BRQ5K4",
      "comment": "",
      "state": {
        "algo": 36199095000000,
        "onl": 2
      }
    },
    {
      "addr": "E36JOZVSZZDXKSERASLAWQE4NU67HC7Q6YDOCG7P7IRRWCPSWXOI245DPA",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "I5Q6RRN44OZWYMX6YLWHBGEVPL7S3GBUCMHZCOOLJ245TONH7PERHJXE4A",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "2GYS272T3W2AP4N2VX5BFBASVNLWN44CNVZVKLWMMVPZPHVJ52SJPPFQ2I",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "D5LSV2UGT4JJNSLJ5XNIF52WP4IHRZN46ZGWH6F4QEF4L2FLDYS6I6R35Y",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "UWMSBIP2CGCGR3GYVUIOW3YOMWEN5A2WRTTBH6Y23KE3MOVFRHNXBP6IOE",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "OF3MKZZ3L5ZN7AZ46K7AXJUI4UWJI3WBRRVNTDKYVZUHZAOBXPVR3DHINE",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "2PPWE36YUMWUVIFTV2A6U4MLZLGROW4GHYIRVHMUCHDH6HCNVPUKPQ53NY",
      "comment": "",
      "state": {
        "algo": 440343426000000,
        "onl": 2
      }
    },
    {
      "addr": "JRGRGRW4HYBNAAHR7KQLLBAGRSPOYY6TRSINKYB3LI5S4AN247TANH5IQY",
      "comment": "",
      "state": {
        "algo": 362684706000000,
        "onl": 2
      }
    },
    {
      "addr": "D7YVVQJXJEFOZYUHJLIJBW3ATCAW46ML62VYRJ3SMGLOHMWYH4OS3KNHTU",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "PZJKH2ILW2YDZNUIYQVJZ2MANRSMK6LCHAFSAPYT6R3L3ZCWKYRDZXRVY4",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "3MODEFJVPGUZH3HDIQ6L2MO3WLJV3FK3XSWKFBHUGZDCHXQMUKD4B7XLMI",
      "comment": "",
      "state": {
        "algo": 130000000000000,
        "onl": 2
      }
    },
    {
      "addr": "WNSA5P6C5IIH2UJPQWJX6FRNPHXY7XZZHOWLSW5ZWHOEHBUW4AD2H6TZGM",
      "comment": "",
      "state": {
        "algo": 130000000000000,
        "onl": 2
      }
    },
    {
      "addr": "OO65J5AIFDS6255WL3AESTUGJD5SUV47RTUDOUGYHEIME327GX7K2BGC6U",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "DM6A24ZWHRZRM2HWXUHAUDSAACO7VKEZAOC2THWDXH4DX5L7LSO3VF2OPU",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "NTJJSFM75RADUOUGOBHZB7IJGO7NLVBWA66EYOOPU67H7LYIXVSPSI7BTA",
      "comment": "",
      "state": {
        "algo": 18099548000000,
        "onl": 2
      }
    },
    {
      "addr": "DAV2AWBBW4HBGIL2Z6AAAWDWRJPTOQD6BSKU2CFXZQCOBFEVFEJ632I2LY",
      "comment": "",
      "state": {
        "algo": 1000000000000,
        "onl": 2
      }
    },
    {
      "addr": "M5VIY6QPSMALVVPVG5LVH35NBMH6XJMXNWKWTARGGTEEQNQ3BHPQGYP5XU",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "WZZLVKMCXJG3ICVZSVOVAGCCN755VHJKZWVSVQ6JPSRQ2H2OSPOOZKW6DQ",
      "comment": "",
      "state": {
        "algo": 45248869000000,
        "onl": 2
      }
    },
    {
      "addr": "XEJLJUZRQOLBHHSOJJUE4IWI3EZOM44P646UDKHS4AV2JW7ZWBWNFGY6BU",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "OGIPDCRJJPNVZ6X6NBQHMTEVKJVF74QHZIXVLABMGUKZWNMEH7MNXZIJ7Q",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "G47R73USFN6FJJQTI3JMYQXO7F6H4LRPBCTTAD5EZWPWY2WCG64AVPCYG4",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "PQ5T65QB564NMIY6HXNYZXTFRSTESUEFIF2C26ZZKIZE6Q4R4XFP5UYYWI",
      "comment": "",
      "state": {
        "algo": 5000000000000,
        "onl": 2
      }
    },
    {
      "addr": "R6S7TRMZCHNQPKP2PGEEJ6WYUKMTURNMM527ZQXABTHFT5GBVMF6AZAL54",
      "comment": "",
      "state": {
        "algo": 1000000000000,
        "onl": 2
      }
    },
    {
      "addr": "36LZKCBDUR5EHJ74Q6UWWNADLVJOHGCPBBQ5UTUM3ILRTQLA6RYYU4PUWQ",
      "comment": "",
      "state": {
        "algo": 5000000000000,
        "onl": 2
      }
    },
    {
      "addr": "JRHPFMSJLU42V75NTGFRQIALCK6RHTEK26QKLWCH2AEEAFNAVEXWDTA5AM",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "64VZVS2LFZXWA5W3S657W36LWGP34B7XLMDIF4ROXBTPADD7SR5WNUUYJE",
      "comment": "",
      "state": {
        "algo": 171945701000000,
        "onl": 2
      }
    },
    {
      "addr": "TXDBSEZPFP2UB6BDNFCHCZBTPONIIQVZGABM4UBRHVAAPR5NE24QBL6H2A",
      "comment": "",
      "state": {
        "algo": 60000000000000,
        "onl": 2
      }
    },
    {
      "addr": "XI5TYT4XPWUHE4AMDDZCCU6M4AP4CAI4VTCMXXUNS46I36O7IYBQ7SL3D4",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "Y6ZPKPXF2QHF6ULYQXVHM7NPI3L76SP6QHJHK7XTNPHNXDEUTJPRKUZBNE",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "6LY2PGUJLCK4Q75JU4IX5VWVJVU22VGJBWPZOFP3752UEBIUBQRNGJWIEA",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "L7AGFNAFJ6Z2FYCX3LXE4ZSERM2VOJF4KPF7OUCMGK6GWFXXDNHZJBEC2E",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "RYXX5U2HMWGTPBG2UDLDT6OXDDRCK2YGL7LFAKYNBLRGZGYEJLRMGYLSVU",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "S263NYHFQWZYLINTBELLMIRMAJX6J5CUMHTECTGGVZUKUN2XY6ND2QBZVY",
      "comment": "",
      "state": {
        "algo": 21647524000000,
        "onl": 2
      }
    },
    {
      "addr": "AERTZIYYGK3Q364M6DXPKSRRNSQITWYEDGAHXC6QXFCF4GPSCCSISAGCBY",
      "comment": "",
      "state": {
        "algo": 19306244000000,
        "onl": 2
      }
    },
    {
      "addr": "34UYPXOJA6WRTWRNH5722LFDLWT23OM2ZZTCFQ62EHQI6MM3AJIAKOWDVQ",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "EDVGNQL6APUFTIGFZHASIEWGJRZNWGIKJE64B72V36IQM2SJPOAG2MJNQE",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    },
    {
      "addr": "RKKLUIIGR75DFWGQOMJB5ZESPT7URDPC7QHGYKM4MAJ4OEL2J5WAQF6Z2Q",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "M4TNVJLDZZFAOH2M24BE7IU72KUX3P6M2D4JN4WZXW7WXH3C5QSHULJOU4",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "WQL6MQS5SPK3CR3XUPYMGOUSCUC5PNW5YQPLGEXGKVRK3KFKSAZ6JK4HXQ",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "36JTK4PKUBJGVCWKXZTAG6VLJRXWZXQVPQQSYODSN6WEGVHOWSVK6O54YU",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "YFOAYI4SNXJR2DBEZ3O6FJOFSEQHWD7TYROCNDWF6VLBGLNJMRRHDXXZUI",
      "comment": "",
      "state": {
        "algo": 30000000000000,
        "onl": 2
      }
    },
    {
      "addr": "XASOPHD3KK3NNI5IF2I7S7U55RGF22SG6OEICVRMCTMMGHT3IBOJG7QWBU",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "H2AUGBLVQFHHFLFEPJ6GGJ7PBQITEN2GE6T7JZCALBKNU7Q52AVJM5HOYU",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "GX3XLHSRMFTADVKJBBQBTZ6BKINW6ZO5JHXWGCWB4CPDNPDQ2PIYN4AVHQ",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "VBJBJ4VC3IHUTLVLWMBON36Y5MPAMPV4DNGW5FQ47GRLPT7JR5PQOUST2E",
      "comment": "",
      "state": {
        "algo": 4524887000000,
        "onl": 2
      }
    },
    {
      "addr": "7AQVTOMB5DJRSUM4LPLVF6PY3Y5EBDF4RZNDIWNW4Z63JYTAQCPQ62IZFE",
      "comment": "",
      "state": {
        "algo": 50000000000000,
        "onl": 2
      }
    },
    {
      "addr": "B4ZIHKD4VYLA4BAFEP7KUHZD7PNWXW4QLCHCNKWRENJ2LYVEOIYA3ZX6IA",
      "comment": "",
      "state": {
        "algo": 40000000000000,
        "onl": 2
      }
    },
    {
      "addr": "G5RGT3EENES7UVIQUHXMJ5APMOGSW6W6RBC534JC6U2TZA4JWC7U27RADE",
      "comment": "",
      "state": {
        "algo": 10000000000000,
        "onl": 2
      }
    },
    {
      "addr": "5AHJFDLAXVINK34IGSI3JA5OVRVMPCWLFEZ6TA4I7XUZ7I6M34Q56DUYIM",
      "comment": "",
      "state": {
        "algo": 20000000000000,
        "onl": 2
      }
    }
  ],
  "fees": "Y76M3MSY6DKBRHBL7C3NNDXGS5IIMQVQVUAB6MP4XEMMGVF2QWNPL226CA",
  "id": "v1.0",
  "network": "mainnet",
  "proto": "https://github.com/algorandfoundation/specs/tree/5615adc36bad610c7f165fa2967f4ecfa75125f0",
  "rwd": "737777777777777777777777777777777777777777777777777UFEJ2CI",
  "timestamp": 1560211200
}`

func GetDefaultMainNetGenesis(ctx context.Context) model.Genesis {
	var g model.Genesis
	if err := json.Unmarshal([]byte(DefaultMainNetGenesis), &g); err != nil {
		runtime.LogErrorf(ctx, "error could not unmarshal default mainnnet genesis json: %s\n", err)
		return g
	}
	return g
}

func LoadMainNetGenesis(ctx context.Context, path string) model.Genesis {
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
