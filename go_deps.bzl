load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "build_buf_gen_go_bufbuild_protovalidate_connectrpc_go",
        importpath = "buf.build/gen/go/bufbuild/protovalidate/connectrpc/go",
        sum = "h1:YUkY57KX9qunqxCHg38Fe4IXk4hCntFPcR4ut61T8cE=",
        version = "v1.16.1-20231115204500-e097f827e652.1",
    )
    go_repository(
        name = "build_buf_gen_go_bufbuild_protovalidate_protocolbuffers_go",
        importpath = "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go",
        sum = "h1:CX+739UtmZMp913WhJ3z2v87wJC5WiBeml6URXygYGw=",
        version = "v1.34.0-20240401165935-b983156c5e99.1",
    )
    go_repository(
        name = "build_buf_gen_go_ride_driver_connectrpc_go",
        importpath = "buf.build/gen/go/ride/driver/connectrpc/go",
        sum = "h1:YV1BApGbTVbROXM7TeC23vUTOuDVume+jNNrVN6jzPc=",
        version = "v1.16.1-20240127091614-32d65fcb4c5c.1",
    )
    go_repository(
        name = "build_buf_gen_go_ride_driver_protocolbuffers_go",
        importpath = "buf.build/gen/go/ride/driver/protocolbuffers/go",
        sum = "h1:WlO2AIofGQR8FGpkJ14t48kR7pjtMkPf86zKDpa7i3Q=",
        version = "v1.34.0-20240127091614-32d65fcb4c5c.1",
    )
    go_repository(
        name = "build_buf_gen_go_ride_payments_connectrpc_go",
        importpath = "buf.build/gen/go/ride/payments/connectrpc/go",
        sum = "h1:IhQnMtigAo3GLhOqUnTCwKl2lxVO+a2FuWMUCKQ8784=",
        version = "v1.16.1-20240224142941-a019682ec9da.1",
    )
    go_repository(
        name = "build_buf_gen_go_ride_payments_protocolbuffers_go",
        importpath = "buf.build/gen/go/ride/payments/protocolbuffers/go",
        sum = "h1:zOcoTN8VbCzbuPM9VoJGXWpSLWie5rxWEaNUXAlZH6g=",
        version = "v1.34.0-20240224142941-a019682ec9da.1",
    )
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
        version = "v0.0.0-20190523083050-ea95bdfd59fc",
    )
    go_repository(
        name = "com_connectrpc_authn",
        importpath = "connectrpc.com/authn",
        sum = "h1:m5weACjLWwgwcjttvUDyTPICJKw74+p2obBVrf8hT9E=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_connectrpc_connect",
        importpath = "connectrpc.com/connect",
        sum = "h1:rOdrK/RTI/7TVnn3JsVxt3n028MlTRwmK5Q4heSpjis=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_github_aidarkhanov_nanoid",
        importpath = "github.com/aidarkhanov/nanoid",
        sum = "h1:yxyJkgsEDFXP7+97vc6JevMcjyb03Zw+/9fqhlVXBXA=",
        version = "v1.0.8",
    )
    go_repository(
        name = "com_github_antlr4_go_antlr_v4",
        importpath = "github.com/antlr4-go/antlr/v4",
        sum = "h1:lxCg3LAv+EUK6t1i0y1V6/SLeUi0eKEKdhQAlS8TVTI=",
        version = "v4.13.0",
    )
    go_repository(
        name = "com_github_bufbuild_protovalidate_go",
        importpath = "github.com/bufbuild/protovalidate-go",
        sum = "h1:U/V3CGF0kPlR12v41rjO4DrYZtLcS4ZONLmWN+rJVCQ=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:9F2/+DoOYIOksmaJFPw1tGFy1eDnIJXg+UHjuD8lTak=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_chromedp_cdproto",
        importpath = "github.com/chromedp/cdproto",
        sum = "h1:aPflPkRFkVwbW6dmcVqfgwp1i+UWGFH6VgR1Jim5Ygc=",
        version = "v0.0.0-20230802225258-3cf4e6d46a89",
    )
    go_repository(
        name = "com_github_chromedp_chromedp",
        importpath = "github.com/chromedp/chromedp",
        sum = "h1:dKtNz4kApb06KuSXoTQIyUC2TrA0fhGDwNZf3bcgfKw=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_chromedp_sysutil",
        importpath = "github.com/chromedp/sysutil",
        sum = "h1:+ZxhTpfpZlmchB58ih/LBHX52ky7w2VhQVKQMucy3Ic=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        importpath = "github.com/chzyer/readline",
        sum = "h1:upd/6fQk4src78LMRzh5vItIt361/o4uq553V8B5sGI=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:WBZRG4aNOuI15bLRrCgN8fCq8E5Xuty6jGbmSNEvSsU=",
        version = "v0.0.0-20191209042840-269d4d468f6f",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:jQCWAUqqlij9Pgj2i/PB79y4KOPYVyFYdROxgaCwdTQ=",
        version = "v0.0.0-20231128003011-0fa0005c9caa",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_dragonfish_go_v2",
        importpath = "github.com/dragonfish/go/v2",
        sum = "h1:uAebil+IwIxqa4keuFdpReeQuGLdI3aOpAEOUjQU/r4=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:4X+VP1GHd1Mhj6IB5mMeGbLCleqxjletLK6K0rbxyZI=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:gVPz/FMfvh57HdSJQyvBtF00j8JU4zdyUgIUNhlgg0A=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig_v3",
        importpath = "github.com/go-task/slim-sprig/v3",
        sum = "h1:sUs3vkvUymDpBKi3qH1YSqBQk9+9D/8M2mN1vB6EwHI=",
        version = "v3.0.0",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:exrUm0f4YX0L7EBwZHuCF4GDp8aJfVeBrlLQrs6NqWU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        importpath = "github.com/gobwas/pool",
        sum = "h1:xfeeEhW7pwmX8nuLVlqbzVc7udMDrwetjEv+TZIz1og=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        importpath = "github.com/gobwas/ws",
        sum = "h1:F2aeBZrm2NDsc7vbovKrWSogd4wvfAxg0FQ89/iqOTk=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:uCdmnmatrKCgMBlM4rMuJZWOkPDqdbZPnrMXDY4gI68=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:7cYmW1XlMY7h7ii7UhUyChSgS5wUJEnm9uZVTGqOWzg=",
        version = "v4.5.0",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v5",
        importpath = "github.com/golang-jwt/jwt/v5",
        sum = "h1:d/ix8ftRUorsN+5eMIlF4T6J8CAt9rch3My2winC1Jw=",
        version = "v5.2.0",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_google_cel_go",
        build_naming_convention = "go_default_library",  #HACK: https://github.com/google/cel-go/issues/801
        importpath = "github.com/google/cel-go",
        sum = "h1:nDx9r8S3L4pE61eDdt8igGj8rf5kjYR3ILxWIpWNi84=",
        version = "v0.20.1",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:OF1IPgv+F4NmqmJ98KTjdN97Vs1JxDPB3vbmYzV2dpk=",
        version = "v0.2.1-0.20230907215043-c6f79328ddf9",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:IqNFLAmvJOgVlpdEBiQbDc2EwKW77amAycfTuWKdfvw=",
        version = "v3.3.2",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:k7nVchz72niMH6YLQNvHSdIE7iqsQxK1P41mySCvssg=",
        version = "v0.0.0-20240424215950-a892ee059fd6",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        sum = "h1:vWQspBTo2nEqTUFita5/KeEWlUL8kQObDFbub/EN9oE=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_google_wire",
        importpath = "github.com/google/wire",
        sum = "h1:HBkoIh4BdSxoyo9PveV8giw7ZsaBOvzWKfcg/6MrVwI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_file_proto_mode = "disable_global",  #HACK: https://github.com/bazelbuild/rules_go/issues/3625
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:5/zPPDvw8Q1SuXjrqrZslrqT7dL/uJT2CQii/cLCKqA=",
        version = "v2.12.3",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:KwWnWVWCNtNq/ewIX7HIKnELmEx2nDP42yskD/pi7QE=",
        version = "v0.0.0-20240312041847-bd984b5ce465",
    )
    go_repository(
        name = "com_github_ilyakaznacheev_cleanenv",
        importpath = "github.com/ilyakaznacheev/cleanenv",
        sum = "h1:0VNZXggJE2OYdXE87bfSSwGxeiGt9moSR2lOrsHHvr4=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        sum = "h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_josharian_intern",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_micahparks_jwkset",
        importpath = "github.com/MicahParks/jwkset",
        sum = "h1:TkMEDKFYAcOgLdl2o6z4mI0EDKK8YN1ip1yRi0bVdbw=",
        version = "v0.5.10",
    )
    go_repository(
        name = "com_github_micahparks_keyfunc",
        importpath = "github.com/MicahParks/keyfunc",
        sum = "h1:lhKd5xrFHLNOWrDc4Tyb/Q1AJ4LCzQ48GVJyVIID3+o=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_micahparks_keyfunc_v3",
        importpath = "github.com/MicahParks/keyfunc/v3",
        sum = "h1:SuFGdd3HvlwEceJvlEEfjJjvOiq69hS0wqM5iMbTlaA=",
        version = "v3.2.4",
    )
    go_repository(
        name = "com_github_mmcloughlin_geohash",
        importpath = "github.com/mmcloughlin/geohash",
        sum = "h1:9w1HchfDfdeLc+jFEf/04D27KP7E2QmpDu52wPbJWRE=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:7eMhcy3GimbsA3hEnVKdw/PQM9XN9krpKVXsZdph0/g=",
        version = "v2.17.2",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:dsYjIxxSR755MDmKVsaFQTE22ChNBcuuTWgkUDSubOk=",
        version = "v1.33.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_github_stoewer_go_strcase",
        importpath = "github.com/stoewer/go-strcase",
        sum = "h1:g0eASXYtp+yvN9fK8sH94oCIk0fau9uV1/ZdJ0AVEzs=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:ZaGT6LiG7dBzi6zNOvVZwacaXlmf3lRqnC4DQzqyRQw=",
        version = "v0.112.2",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:vO95gvBi7qUgfA9SflexQs9hB4U4tnri/GwADIrLQy8=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:GgdNoDwZR5RIO3j8XwXqa6Gc6q5mP3KYMdFC7FEVyG4=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:YWeqD4BjYwrmY4fa+isGcw0P81lJ3dKVxbWxdBchoiU=",
        version = "v1.67.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:O0fj88npvQFxg8LfXo7fArcSrC/wtAstGuWQ7dCHWjg=",
        version = "v0.23.2",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:DO5Vn3zmY1aDyfoqni8e8+x+lwrfLCoAAbEui9NB0y8=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:z08Xuv7ZtaB2d4jsJi9/WhbnnI5s19wlLDZpssn3Fus=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:o1C/+IvzwYeV1doum61XmJQ/Bwpk/4+2DT1JyVu2x64=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:qYrjEHEFY7+CL4QlHIHuwTgrTnZbSKzdPFqgjZDsQNo=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:sUrR96yokdL6tTTXK0X13V1TLMta8/1u328bRG5lWZc=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:SSvoD0ofOydm5gA1++15pW9VPgQbk0OmNlcb7JczoO4=",
        version = "v1.14.9",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:mCqyoaDjDzaW1RqmmQtCJuawb9nca5bEu7HvVcpZDwg=",
        version = "v1.19.1",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:xieyFA+JKyTDkO/Z9UyVEpkHW8pDYykU51O4G0pvXEg=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_auth",
        importpath = "cloud.google.com/go/auth",
        sum = "h1:PRyzEpGfx/Z9e8+lHsbkoUVXD0gnu4MNmm7Gp8TQNIs=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_auth_oauth2adapt",
        importpath = "cloud.google.com/go/auth/oauth2adapt",
        sum = "h1:+TTV8aXpjeChS9M+aTtN/TjdQnzJvmzKFt//oWu7HX4=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:w9AyogtMLXbcy5kzXPvk/Q3MGQkgJH7ZDB8fAUUxTt8=",
        version = "v1.13.7",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:W4oSMS6vRCo9DLr1RPyDP8oeLverbvhJRzaZSsipft8=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:i8shmhASiPJ/DFHhqwBcNQep+uFlHt53Txal/wsc9ko=",
        version = "v1.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:KBcujO3QRvBIwzZLtvQEPB9SXdovHnMBx0V/uhucH9o=",
        version = "v1.0.6",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:w2Goy9n6gh91LVi6B2Sc+HpBl8WbWhIyzdvVvrAuEIw=",
        version = "v1.61.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:GbOg1uGvoV8FXxMStFoNcq5z9AEUwCpKt/6GNcuDSZM=",
        version = "v1.18.5",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:RHnEM4HXbWShlGhPA0Jzj2YYETCHxmisNMU0OE2fXQM=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:XURrQhj5COWAEvICivbGID/Hu67AvMYHAhMRIyc3Ux8=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:PrplNaAS6Dn187e+OcGzyEKETX8iL3tCaDqcPPW7Zoo=",
        version = "v1.17.7",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:zkCG1dBezxRM3dtgQ9h1Y+IJ7V+lARWgp0l9k/SZsfU=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:Q47KKoA0zsNcC9U5aCmop5TPPItVq4cx7Wwqgra+5PU=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:Y0HUuiCAVk9BojLItOycBl91tY25NXH8oFsyi1IC/U4=",
        version = "v1.12.8",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:uHf0NN2nvxl1Gh4QO83yRCOdMK4zivtMS5gv0dEX0hg=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:Tz+eQXMEqDIKRsmY3cHTL6FVaynIjX2QxYC4trgAKZc=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:46ertIh+cGkTg/lN7fN+TOx09SoM65dpdUp96vXBcMY=",
        version = "v1.13.2",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:Vbu/3PZNrgV1Z5DGcRubQdUccX/uMUDNc+NgHNIfbEk=",
        version = "v1.35.1",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:mSrneOVadcpnDZHJebg+ts/10azGTUKOCSQET7KdT7g=",
        version = "v0.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:czcba5mxwRM5V//jSadyig0y+8aOHmN7gUl9GbHu59E=",
        version = "v1.20.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:wKEakCbRevlwsWqTn34pWJUFmdbx0HKwpRH6HhU7NIs=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:MiK1Us7YP9+sdNViUE4X2B2vLScrKcjOPw5b6uamZvE=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:ViFnMnUK7LNcWvisZgihxXit76JxSHFeijYI5U/gjOE=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:M6irSHns6VxMro+IbvDxDJLD6tkfjlW+mo2MPaM23KA=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:x5A/rDYv4ZrGsjona+DRx05+pHjRzyyGwr+Qqg5Xitk=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:RNMG5ffWKdbWOkwvjC4GqxLaxEaWFpm2hQCF2WFW/vo=",
        version = "v2.4.2",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:qM60MGNTGsSJuzAziVJjtRA7pGby2dA8OuqdVRe/lYo=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:LrZmu9l/qjoX/ilR+ECSMyO6tDYpijp3RR5LBM0HjpU=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:FfNUy9j3aRQ99L4a5Rdm82RMuiw0BIe3lpPn2ykom8k=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:ad9mYYDLTPjLcEIlvwsFJQxmyXZx2sv1DN5sjkl2hTY=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:C9wQ0odRYQsar0XqwCQb0c13BkRBsoSjOaejOg5ntgQ=",
        version = "v1.53.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:5SzyGs784ql/QS5ktXpUEbKiWarz4pmfjDe/A6WkjeQ=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:tLn+VjEf+xBhNo+UpecHFsrnx4RB2AQP2WH1DvggBUQ=",
        version = "v1.28.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:IixFIMRzUJWZUAOe8s/K2X4Bvtp0A3xjHLljfNC4aSo=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:xa6MIQhGylE24QdWaxhfIfAJE3Pupcr+i77WEx3NJrg=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:p5Y7ZNVPiV9pEAHzvWiPcSiQRMQqcuHxOP0ZOP0vVww=",
        version = "v1.6.8",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:we+qx5uCZ88aQzQS3MJXRvAh/ik+EmqVyjcW1oYFW44=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:CpRnsUpMU5gxUKyfh7TD0SM+E+7E4ORaDea2JctKfpY=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:/k8ppuWOtNuDHt2tsRV42yI21uaGnKDEQnRFeBpbFF8=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:83bd2lCgtu2nLbX2jrqsrQhIs7VuVA1N6Op5syeRVIg=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:8WsRfKbElUgN8NTD1sYfRDteztxl9KPRvx4SgSGGXcg=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:BfXsTXYs5xlicAlgbtlo8Cw+YdzU3PrlBg7dATJUwrk=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:bHwcvgh8AmcYm6p6/ZrWW3a7J7sKBDtqtsyVXKssnPs=",
        version = "v0.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:YtLvUpB/YGjrrloG6IXAo45BdviOHr4Emde3ABYuhlI=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:06Jg3JeLslEfBYX1sDqOPLnF7a3wmhNcDUXF/fVOb50=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:r7umDwhj+BQyz0ScZMp4QrGXjSTI3ZINnpgU2nlB/K0=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:rcuRS9XfOgr1v6TAoihVeSXntOnpVhFlVHtPfgOkLAo=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:wtd+r415yrfZ8LsB6yH6WrOZ26tYt7w6wy3i5a4HQZ8=",
        version = "v1.4.7",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:M9SKIj9eoxoXCzytkLZVAuf5wmoui1OeDqEjC97wRbY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:ouZjTxCqDNEdxWfaAAbRzG22s/2iewRw6JPARQL+0vc=",
        version = "v1.15.9",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:kOYJEcuZgyUX/i/4DFrfXPcrddm1XCQD2lDI5hIFmZQ=",
        version = "v1.12.5",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:qqEmApr5YFOQjkrU8Jy6o6QpkESqfGbfrE6bnUZZbV8=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:iEIOXFO9EmSiTjDmfpbRjOxECO7R8C7b8IXUGOj7xZw=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:WLbHekDbjK1fVFD3ibpFFVoyizlLRl73I7YKuAKilhU=",
        version = "v0.5.7",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:uWA9WQyfA0JdkeAFymWUsa3qE9tC33LUElla790Ou1A=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:FHj/3v2rGBOP1ktPAm4IsM4T0l6L8Hhh9/TenXJpvlE=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:izgww3TlyvWyDWdFKnrASpbh12IkAuw8o2ION8sAjX0=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:hE7f3ze3+eWh/EbYXEz7oXkm0LXcr7UCoLklwi7gsLU=",
        version = "v1.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:otHcJkci5f/sNRedrSM7eM81QRnu0yZ3HvkvWGphABA=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:NCXf8hfQi+Kmr56QJezXRZ6GPb80ZI7El1XztyUuLQI=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:jYpQ86mZ7OYZc7WadvCIlIaPXmXhr5nD7wgE/ekMVpM=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:Ex1/aYkA0areleSmOGXHvEFBGohteIYJr2SGPrjOUe0=",
        version = "v1.13.2",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:aepEkfiwOvUL9eu3ginVZhTaXDRHncQKi9lTT1BycH0=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:sFU1ETg1HfIN/Tev8gD0dleAITLv7cHp0JClwFmJ6bo=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:FPfowA/LEckKTQT0A4NJMI2bSou999c2ZyFX1zGiYxY=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:C2WL4ZnclXsh4XickGhKYKlPjqVZj35y1sbRjdsZ3g4=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:fGftW2bPi8vTjQm57xlwtLBZQcrgC+c3HMFBzJ+KWPc=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:HXsXGFaFaLTklwKgSob/GSE+c3verYDQDgreFaosxyc=",
        version = "v1.12.7",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:7AgOWH1oMPrB1AVU0/f47ADdOt+XfdBY7QRb8tcMUp8=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:CbCjfR/pgDHyRMu94o9nuGwaONEcarWnUfSGGw+I2ZI=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:LGt85MZUKlq9oqsbBL9+M6jAyeuR1TtCx6k5HfAQxTY=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:wGZKKJhYyuf4gcAEywQqQ6F19yxhBJGnzgyxOTbJjBw=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:0uEEfaB1VIJzabPpwpZf44zWAKAme3zwKKxHk7vJQxQ=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:+QG02kE63W13vXI+rwAxFF3EhGX6K7gXwFz9OKwKcHw=",
        version = "v2.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:N6n/TEr0FQzeP4ZtvF5daMszOhdZI94uMiPiAi9kFMo=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:v9x75vXP5wMXw3QG3xmgjVHLlqYufuLn/ht3oNWCA3w=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:vXRKu2ekEBqiTSdm4Qu39MqMob0PP6IQwPyQUpEgue4=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:SdvD0PaPX60+yeKoSe16mawFpM0EPuiPPihTIVlhRsY=",
        version = "v1.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:88SlpWtogkwjMuYTEl//qm36azX1OpawThAyvXT/hHw=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:msP5a8BOxVym2DvoubeWAxAeV6VhYkKnYHc2XOkd/+U=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:E4Z5e681Qh7UJrJRMCgYhp+3tkcoXiaKGh3UZmUPaAQ=",
        version = "v1.3.7",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:Jn/unfNUgRiNJRc1nrApzimKiVj91UYlLT8mMfpUu48=",
        version = "v1.10.8",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:nQ/Ca2Gzm/OEP8tr1hiFdHRi5wAnAmsm9qTjwkivyrQ=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:9Jn8BJpkq8MflNzTdrX4m+SVp2+WeqVhbFiwyNIoXuM=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:Y8C0I/mzLbaxAl5cw3EaLox0Rvpy+VUwEuCGWIQDMU8=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:sWvEqg3CLcRu2dgGnQ4479CzGFgGJxlcGtRtlPL450M=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:HxCzcUxSsCh6FJWkmbOUrGI1sKe4E1Yy4vaykn4RhJ4=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:P7XRZDjBnNw+3tHkPrtWzcxtC3Cqhm+X0vWrO61Ry58=",
        version = "v1.61.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:TcWEAOLQH1Lb2fhHS6/GjvAh+ue0dt4xUDHXHG6vF04=",
        version = "v1.23.1",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:VEpDQV5CJxFmJ6ueWNsKxcr1QAYOXEgxDa+sBbJahPw=",
        version = "v1.40.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:CXmoNEvz7y2NtHFZuH3Z8ASN43rxRINWa2Q/IlBzM2k=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:RoyEtftfJrbwJcu63zuWE4IjC76xMyVsJBhmleIp3bE=",
        version = "v1.6.8",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:qR6Mu+EM2OfaZR1/Rl8BDBTVfi2X5OtwKKvJRSQyG+o=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:ngQokxUB1z2gvHn3vAf04m7SFnNYMiQIIpny81fCGAs=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:gK8z2BIJQ3KIYGddw9RJLne5Fx0FEXkrEQzPaeEYVvk=",
        version = "v1.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:g+B29z4gtRGsiKDoTF+bNeH25bLRokAaElygX2FcZkE=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:YZHUgiIHE77SdZNT6gjagoni5GRcUkkrm4YCFFPqtBw=",
        version = "v1.20.6",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:SKBkFTuOclESLjQL1LwraqVFm2fL5oL9tbzKITU+FOY=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:j9RxG8DcyJO/D7/ps2pOey8VZys+TMqF79bWAhuM7QU=",
        version = "v2.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:bf2qKqEN7iqT62IptQ/FDadoDLJI9sthyrW3PVaH8bY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:x4KwHB4JlBEzMaITVhrbbpHrU+2I5LrlvHGEEluT0vc=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:F5woMLufKnshmDvPVxCzoC+Di12RYXQ1W8kNmpBT8z0=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:EWTSVagWWeQjVAsebiF/wJMwC5bq6Zz3LqOmD9Uid4s=",
        version = "v1.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:R5OW5SNRqD0DSEmyWLUMNYAXWYnz/NLSXBawVFrc9a0=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:2bE69mh68law1UZWPjgmvOQsjsGSppRudABAXwNAy58=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_firebase_go_v4",
        importpath = "firebase.google.com/go/v4",
        sum = "h1:Tc9jWzMUApUFUA5UUx/HcBeZ+LPjlhG2vNRfWJrcMwU=",
        version = "v4.14.0",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:YR8cESwS4TdDjEe65xsg0ogRM/Nc3DYOhEAlW+xobZo=",
        version = "v1.0.0-20190902080502-41f04d3bba15",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_olympos_encoding_edn",
        importpath = "olympos.io/encoding/edn",
        sum = "h1:slmdOY3vp8a7KQbHkL+FLbvbkgMqmXojpFUO/jENuqQ=",
        version = "v0.0.0-20201019073823-d3554ca0b0a3",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:4Pp6oUg3+e/6M4C0A/3kJ2VYa++dsWVTtGgLVj5xtHg=",
        version = "v0.49.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:jq9TW8u3so/bN+JPT166wjOI6/vQPF6Xe7nMNIltagk=",
        version = "v0.49.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:0LAOdjNmQeSTzGBzduGe/rU4tZhMwL5rWgtp9Ku5Jfo=",
        version = "v1.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:6EhoGWWK28x1fbpA4tYTOWBkPefTDQnb8WSGXlc88kI=",
        version = "v1.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:YMPPDNymmQN3ZgczicBY3B6sf9n62Dlj9pWD3ucgoDw=",
        version = "v1.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:CsKnnL4dUAr/0llH9FKuc698G04IrpWV0MQA/Y1YELI=",
        version = "v1.24.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:8a0p/BbPa65GlqGWtUKxot4p0TV8OGOfyTjtmkXNXmk=",
        version = "v0.177.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_appengine_v2",
        importpath = "google.golang.org/appengine/v2",
        sum = "h1:MSqyWy2shDLwG7chbwBJ5uMyw6SNqJzhJHNDwYB0Akk=",
        version = "v2.0.2",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:HjgkYCl6cWQEKSHkpUp4Q8VB74swzyBwTz1wtTzahm0=",
        version = "v0.0.0-20240506185236-b8a5c65736ae",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:AH34z6WAGVNkllnKs5raNq3yRq93VnjBG6rpfub/jYk=",
        version = "v0.0.0-20240506185236-b8a5c65736ae",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:GtsRfMHDREQPg/snOM0QudeC54kX7UqodmmK4uELHLQ=",
        version = "v0.0.0-20240429193739-8cf5692501f6",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:DujSIu+2tC9Ht0aPNA7jgj23Iq8Ewi5sgkQ++wdvonE=",
        version = "v0.0.0-20240429193739-8cf5692501f6",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:MUeiw1B2maTVZthpU5xvASfTh3LDbxHd6IJ6QQVU+xM=",
        version = "v1.63.2",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        build_file_proto_mode = "disable_global",  # Manually added to fix build. See https://github.com/golang/protobuf/issues/1611
        sum = "h1:9ddQBjfCyZPOHPUiPxpYESBLc+T8P3E+Vo4IbKZgFWg=",
        version = "v1.34.1",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:dIJU/v2J8Mdglj/8rJ6UUOM3Zc9zLZxVZwwxMooUSAI=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:aAcj0Da7eBAtrTp03QXWvm88pSyOt+UgdZw2BFZ+lEw=",
        version = "v0.0.0-20240325151524-a685a6edb6d8",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
        version = "v0.0.0-20190313153728-d0100b6bd8b3",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:d/OCCoBEUq33pjydKrGQhw7IlUPI2Oylr+8qLx49kac=",
        version = "v0.25.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:9+E/EZBCbTLNrbN35fHv/a/d/mOBatymz1zbtQrXpIg=",
        version = "v0.19.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:Od9JTbYCk261bKm4M/mw7AklTlFYIa0bIp9BgSm1S8Y=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_telemetry",
        importpath = "golang.org/x/telemetry",
        sum = "h1:IRJeR9r1pYWsHKTRe/IInb7lYvbBVIqOgsX/u0mbOWY=",
        version = "v0.0.0-20240228155512-f48c80bd79b2",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:VnkxpohqXaOBYJtBmEppKUG6mXpi+4O6purfc2+sMhw=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:h1V/4gjBv8v9cjcR6+AR5+/cIYK5N/WAgiv4xlsEtAk=",
        version = "v0.15.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:hz/CVckiOxybQvFw6h7b/q80NTr9IUQb4s1IIzW7KNY=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:+cNy6SZtPcJQH3LJVLOSmiC7MMxXNOb3PU/VUEz+EhU=",
        version = "v0.0.0-20231012003039-104605ab7028",
    )
    go_repository(
        name = "org_uber_go_goleak",
        importpath = "go.uber.org/goleak",
        sum = "h1:xqgm/S+aQvhWFTtR0XK3Jvg7z8kGV8P4X14IzwN3Eqk=",
        version = "v1.2.0",
    )
    go_repository(
        name = "org_uber_go_mock",
        importpath = "go.uber.org/mock",
        sum = "h1:VcM4ZOtdbR4f6VXfiOpwpVJDL6lCReaZ6mw31wqh7KU=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:sI7k6L95XOKS281NhVKOFCUNIvv9e0w4BF8N3u+tCRo=",
        version = "v1.26.0",
    )
