/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package aws

import (
	"hcm/pkg/adaptor/types/image"
	"hcm/pkg/kit"
	"hcm/pkg/logs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// TODO: 临时方案，之后切换更合理的方案
var regionImageIDMap = map[string][]*string{
	"ap-southeast-2": {
		aws.String("ami-0d9f286195031c3d9"),
		aws.String("ami-00149ba676b864e3c"),
		aws.String("ami-0dc5d846beb5d669b"),
		aws.String("ami-0fec995d2c2cd0d4c"),
		aws.String("ami-076b75d50e423acfe"),
		aws.String("ami-0cbf852f55ea404a6"),
		aws.String("ami-07dba8e4546543211"),
		aws.String("ami-05f998315cca9bfe3"),
		aws.String("ami-0a9cfa40cedba2d5a"),
		aws.String("ami-09cf24ffd6d332930"),
		aws.String("ami-03d0155c1ef44f68a"),
		aws.String("ami-0e6b0944740fe3648"),
		aws.String("ami-0a55150d12bfd9900"),
		aws.String("ami-04ce798621838c3b7"),
		aws.String("ami-0f069fc2c8b8aa489"),
		aws.String("ami-0f6917ff067760030"),
		aws.String("ami-06f1687bdcbddb2e1"),
		aws.String("ami-0d00cfb25d8758630"),
		aws.String("ami-0f8ceec51ba3a3a21"),
		aws.String("ami-083987c64cfee7e4a"),
		aws.String("ami-0ae7767061c20b57f"),
		aws.String("ami-0864086ece51b0aea"),
		aws.String("ami-027555f0a9943ed8d"),
		aws.String("ami-05f0436dc5c2d6086"),
		aws.String("ami-0fd8e61e75cd8d4b8"),
		aws.String("ami-0786dab6f7c6de3fd"),
		aws.String("ami-017007b46f8acefcc"),
		aws.String("ami-03e13ad6ac11e6707"),
		aws.String("ami-04a0ab226af889a7a"),
		aws.String("ami-01403b49e8c7d040c"),
		aws.String("ami-03b445535efae04f3"),
		aws.String("ami-06d2baabc0d44fec4"),
		aws.String("ami-09cdc480695e9707a"),
		aws.String("ami-0c371ecb9078a8020"),
		aws.String("ami-0dd6f26dbce3c7408"),
		aws.String("ami-0da5592b19659bc6f"),
		aws.String("ami-08aa884d56ac818c5"),
		aws.String("ami-01e9d90ef3eed8a11"),
		aws.String("ami-04ce798621838c3b7"),
		aws.String("ami-09fb3a51968a858cd"),
		aws.String("ami-0c8e2720d3a97d3ae"),
		aws.String("ami-0c6de7e6b9756b4d2"),
		aws.String("ami-0d37dd644530396ea"),
		aws.String("ami-00e0f5accb5593178"),
		aws.String("ami-03d7d7615da681d05"),
		aws.String("ami-090787f9264a09492"),
		aws.String("ami-0b38a536a525440ba"),
		aws.String("ami-0b229fab078b3ed3a"),
	},
	"ap-northeast-1": {
		aws.String("ami-079a2a9ac6ed876fc"),
		aws.String("ami-031486c2cf674ac5b"),
		aws.String("ami-06560e4f1897491ed"),
		aws.String("ami-0d80f992174a092f1"),
		aws.String("ami-0cd0830ef4d2de449"),
		aws.String("ami-01c18b1f46deb3457"),
		aws.String("ami-0b168f9cd578fe5d5"),
		aws.String("ami-0dcb1703d648a942f"),
		aws.String("ami-0d979355d03fa2522"),
		aws.String("ami-01a777eb1a2618535"),
		aws.String("ami-0d0c6a887ce442603"),
		aws.String("ami-0e3b5beb11b0c9d8e"),
		aws.String("ami-0334f027111424f54"),
		aws.String("ami-00f40208bd8150952"),
		aws.String("ami-07575b952521e30c9"),
		aws.String("ami-0b14237264b822e4d"),
		aws.String("ami-0149b91ac885b5a2e"),
		aws.String("ami-086dc06b7631fc947"),
		aws.String("ami-0f831738506a1f013"),
		aws.String("ami-0ee87a1a729fecb57"),
		aws.String("ami-0b709072dd76c326e"),
		aws.String("ami-035eea5e9cd788b01"),
		aws.String("ami-0cf984fb3688b2be7"),
		aws.String("ami-0e6081ccf47667751"),
		aws.String("ami-08c7aa2ae986f3fb5"),
		aws.String("ami-010e9dd22c48963db"),
		aws.String("ami-0da9188dd3cfc3e92"),
		aws.String("ami-0df2ab489f87b783d"),
		aws.String("ami-04c5e5386d38552c1"),
		aws.String("ami-0d494d6c4a8ba437a"),
		aws.String("ami-0a11ac8aadbdd3f70"),
		aws.String("ami-09732b80267f793a3"),
		aws.String("ami-077775063fb239e88"),
		aws.String("ami-061fc9ed0314bd451"),
		aws.String("ami-02b058454fc973e0c"),
		aws.String("ami-0ff1be38c5cc44c15"),
		aws.String("ami-0f3b9fa47bd1a7a99"),
		aws.String("ami-0a6176315dc755eaa"),
		aws.String("ami-0d5142f63c808d143"),
		aws.String("ami-05eeede9f4255e28e"),
		aws.String("ami-02123f1fc332954db"),
		aws.String("ami-06e9f5b7af2dc6304"),
		aws.String("ami-094e28968a2e5dcd2"),
		aws.String("ami-01e442590cadbb4a6"),
		aws.String("ami-0db6b798d458e3f51"),
		aws.String("ami-064d2a8f528240c59"),
		aws.String("ami-0be18eba4e877b995"),
		aws.String("ami-0602d6048524038f1"),
	},
	"ca-central-1": {
		aws.String("ami-0289605e6e2e851bf"),
		aws.String("ami-03e67b38071653c20"),
		aws.String("ami-0b533530f1b074e0f"),
		aws.String("ami-0b62e7852f992695c"),
		aws.String("ami-097bd6037de54b1dc"),
		aws.String("ami-0b0a407934c566700"),
		aws.String("ami-0abc4c35ba4c005ca"),
		aws.String("ami-076dec802f67b0617"),
		aws.String("ami-034dc61db286d7891"),
		aws.String("ami-0213c7d78ef9e3968"),
		aws.String("ami-09346c22e58ec5479"),
		aws.String("ami-08ae89bc395ec8e13"),
		aws.String("ami-055f614fa03fcdb19"),
		aws.String("ami-0fe1d19db1d0a7026"),
		aws.String("ami-0e8b073dac337b89d"),
		aws.String("ami-034b7b852c189d4c8"),
		aws.String("ami-094fdfed4257c0cde"),
		aws.String("ami-08413dce74940e624"),
		aws.String("ami-04ec26489a8cc5596"),
		aws.String("ami-014561f361af7f5d2"),
		aws.String("ami-0619bb721ee7a36b2"),
		aws.String("ami-01a52cbe9def91b59"),
		aws.String("ami-05402da33f3054f84"),
		aws.String("ami-09bf16f0c9c205552"),
		aws.String("ami-08f7f1cb098ae13fd"),
		aws.String("ami-0253c670135bdef73"),
		aws.String("ami-0a101cebca202bd0e"),
		aws.String("ami-0d06c60138dfa5c64"),
		aws.String("ami-0909e52a67dded531"),
		aws.String("ami-0a7735ecdcee38b31"),
		aws.String("ami-01c7ecac079939e18"),
		aws.String("ami-0b80499bbf1827661"),
		aws.String("ami-073bc9072be4fc8a6"),
		aws.String("ami-0a607e3c45d48ee8f"),
		aws.String("ami-0721ce52e3bbbdb2a"),
		aws.String("ami-04609b5b156500b17"),
		aws.String("ami-08800aca944d1b71d"),
		aws.String("ami-0eb43817799f06e40"),
		aws.String("ami-0aa8e267e0b567cfe"),
		aws.String("ami-06ab5224e8fe4fe85"),
		aws.String("ami-0c54639d4a96f6616"),
		aws.String("ami-0ef3a409d934752c5"),
		aws.String("ami-018e2501f03c4edf8"),
		aws.String("ami-06c60fdb62907a675"),
		aws.String("ami-0e9eb7d822ffe7b3d"),
		aws.String("ami-0335be6b3a84810b7"),
		aws.String("ami-031df0f05feb1fb53"),
		aws.String("ami-02a7993ffba13bc28"),
	},
	"eu-central-1": {
		aws.String("ami-0fa03365cde71e0ab"),
		aws.String("ami-0750be70a912aa1e9"),
		aws.String("ami-08f54b258788948e1"),
		aws.String("ami-0b17e9c74dda27133"),
		aws.String("ami-0453f67283fbdec39"),
		aws.String("ami-0a830948c354b21a5"),
		aws.String("ami-00f136ac48acf2618"),
		aws.String("ami-0ec7f9846da6b0f61"),
		aws.String("ami-07625524674f7c390"),
		aws.String("ami-0cf9380844da84d7e"),
		aws.String("ami-0af956f7bb5f3cfc4"),
		aws.String("ami-0bb82ebfb4046aa30"),
		aws.String("ami-02d0d633d425a6259"),
		aws.String("ami-05298694e0c36ef04"),
		aws.String("ami-01860e0d17bcfee5d"),
		aws.String("ami-0a2329b9fe3dff9db"),
		aws.String("ami-0a6a4562de19e15c2"),
		aws.String("ami-08f13e5792295e1b2"),
		aws.String("ami-0216ffa50f321442e"),
		aws.String("ami-0e0e61fd3601f7ac8"),
		aws.String("ami-094bfbd6ff1bee505"),
		aws.String("ami-07680b615cc2ce1dd"),
		aws.String("ami-04cd8ecda2d249fd4"),
		aws.String("ami-02ea3c6ca28824141"),
		aws.String("ami-0b8bb746896ddb3a8"),
		aws.String("ami-0ac9e212ac8852729"),
		aws.String("ami-0d497a49e7d359666"),
		aws.String("ami-025c7096bc4cac207"),
		aws.String("ami-01d4e68cf9c107037"),
		aws.String("ami-0779aec1a8c186d3a"),
		aws.String("ami-0de2ba495bb64aa31"),
		aws.String("ami-05782ba31e5abaaef"),
		aws.String("ami-0adbf05649665315c"),
		aws.String("ami-0c94a2b1a05f52579"),
		aws.String("ami-04f13840aeeb2ecbc"),
		aws.String("ami-07e14e673b9d75bf6"),
		aws.String("ami-0087bd2f5c26af5ca"),
		aws.String("ami-074ea14c08effb2d8"),
		aws.String("ami-01277f063a1810558"),
		aws.String("ami-0ec0e5af4387e53d0"),
		aws.String("ami-094950f08c57b4f62"),
		aws.String("ami-049fbc5d92e0102fd"),
		aws.String("ami-0319c3c5392769609"),
		aws.String("ami-0f0ff99f194530aa4"),
		aws.String("ami-03e5f555f0e8d030a"),
		aws.String("ami-0876b09ff1b2c153a"),
		aws.String("ami-0c4f0d0904a037ec7"),
		aws.String("ami-03e79a5285159bef6"),
		aws.String("ami-0cf2fff33e5684d9a"),
	},
	"eu-west-1": {
		aws.String("ami-09dd5f12915cfb387"),
		aws.String("ami-0de2a2552e7fe4b4d"),
		aws.String("ami-04b8ac8af30d436d9"),
		aws.String("ami-095671dc7a5de7c35"),
		aws.String("ami-0b04ce5d876a9ba29"),
		aws.String("ami-062e673cc4273dad8"),
		aws.String("ami-017d70213b2dbe508"),
		aws.String("ami-09ee771fad415a6d7"),
		aws.String("ami-0b2dd5cb6105dcd18"),
		aws.String("ami-00aa9d3df94c6c354"),
		aws.String("ami-0e6902f007857ad9c"),
		aws.String("ami-05147510eb2885c80"),
		aws.String("ami-013607984ea916121"),
		aws.String("ami-0c06c45b5455eb326"),
		aws.String("ami-0ca6075c7cf938d2d"),
		aws.String("ami-056b9e344b66fd4b3"),
		aws.String("ami-0c0fd38623a3f80ea"),
		aws.String("ami-0151488636cf5c1e7"),
		aws.String("ami-0f5a46d734b5f2fd3"),
		aws.String("ami-089f338f3a2e69431"),
		aws.String("ami-0083086bc808b1d71"),
		aws.String("ami-088d62fd83a995ab9"),
		aws.String("ami-05f988034b21b4f79"),
		aws.String("ami-0a3a2a49bf7baac68"),
		aws.String("ami-01a0cf264428bf248"),
		aws.String("ami-092ec6fb3154fdb73"),
		aws.String("ami-020375c819f32a382"),
		aws.String("ami-0c4b5c941059e9b4f"),
		aws.String("ami-07c0febd9a263f0c7"),
		aws.String("ami-09545f35b10d3a81e"),
		aws.String("ami-086c7de4a0d025b35"),
		aws.String("ami-0044d1d9a6d25bc11"),
		aws.String("ami-0775f609ab9dfedfc"),
		aws.String("ami-0467ce267e9435b13"),
		aws.String("ami-00170aefee44312e5"),
		aws.String("ami-0e42eb8d6bc4c323d"),
		aws.String("ami-09eeea9162a4f6629"),
		aws.String("ami-091780e2a3eb42d46"),
		aws.String("ami-0a03160df3d66257b"),
		aws.String("ami-08456538e3a727106"),
		aws.String("ami-056b9e344b66fd4b3"),
		aws.String("ami-0ac03fb170870f7c7"),
		aws.String("ami-0f20f820a86c541d3"),
		aws.String("ami-049798e7835bf34ea"),
		aws.String("ami-0c9e39a9c031878a2"),
		aws.String("ami-0c0679dd6ee951f93"),
		aws.String("ami-0fadca9e429399a30"),
		aws.String("ami-04b82270e2c61ea45"),
		aws.String("ami-0300040ac527987d6"),
		aws.String("ami-0a2f6bffd5f5ec6f9"),
	},
	"eu-west-2": {
		aws.String("ami-0cd8ad123effa531a"),
		aws.String("ami-06a566ca43e14780d"),
		aws.String("ami-08d9bb4bfe39be5c2"),
		aws.String("ami-0952d71efe9366064"),
		aws.String("ami-09744628bed84e434"),
		aws.String("ami-073bb7464cc51df7c"),
		aws.String("ami-048191dc593b5cb31"),
		aws.String("ami-0ddb10e73cf07b977"),
		aws.String("ami-068b5b133ab6ab6d9"),
		aws.String("ami-0b3053411345882ee"),
		aws.String("ami-085d666b72d18c24c"),
		aws.String("ami-0221ff515cf4a0f21"),
		aws.String("ami-04265a7e5895f5c58"),
		aws.String("ami-0d93d81bb4899d4cf"),
		aws.String("ami-062533e1912bab905"),
		aws.String("ami-0a7122e9631d99380"),
		aws.String("ami-08e20fcc8b3af77d1"),
		aws.String("ami-08b6430a61cac945f"),
		aws.String("ami-0c5c78e202de8e5a9"),
		aws.String("ami-03938037a784195d3"),
		aws.String("ami-0a53d4675d3c3bfd1"),
		aws.String("ami-04e4fab13770ee2d7"),
		aws.String("ami-0b9e632db6d766cf6"),
		aws.String("ami-08158c4707cf369ad"),
		aws.String("ami-0784a9180b772a730"),
		aws.String("ami-028a5cd4ffd2ee495"),
		aws.String("ami-01246f7648fa84bb6"),
		aws.String("ami-054df4a1d62dbc444"),
		aws.String("ami-070ee3d06fc5893d7"),
		aws.String("ami-0ef262972e641bb3e"),
		aws.String("ami-0626a64ba55cc938c"),
		aws.String("ami-05f4add752ba96178"),
		aws.String("ami-0b97b24224d89b8ff"),
		aws.String("ami-0fc7889c83753a456"),
		aws.String("ami-05677b62c0c528109"),
		aws.String("ami-01803cab1edc48df7"),
		aws.String("ami-0a0953102a1ac986f"),
		aws.String("ami-03831eb3f00dc1bac"),
		aws.String("ami-089af9a3254595f46"),
		aws.String("ami-0c8219c459ea82b06"),
		aws.String("ami-01a47097cc1a934e6"),
		aws.String("ami-03450cbc86322efab"),
		aws.String("ami-09e606a96109408a5"),
		aws.String("ami-0e2cb5b1e923994c9"),
		aws.String("ami-066297ef75f7bc830"),
		aws.String("ami-0ba1bec9bb08e6015"),
		aws.String("ami-0582d5cf3037b773a"),
	},
	"eu-west-3": {
		aws.String("ami-069fa606c9a99d947"),
		aws.String("ami-09e6612d6af6625c7"),
		aws.String("ami-01b305bdc62291ce1"),
		aws.String("ami-05e8e219ac7e82eba"),
		aws.String("ami-01a3ab628b8168507"),
		aws.String("ami-01daacddaafcdd876"),
		aws.String("ami-07b51a8e58bf6755d"),
		aws.String("ami-0fdf5a3be18a0a653"),
		aws.String("ami-05f184987999daffd"),
		aws.String("ami-06ee5ff817eed44ba"),
		aws.String("ami-0f13ee4d0a693dadd"),
		aws.String("ami-0b658662f2f2bf294"),
		aws.String("ami-07649b7bf99ca07b8"),
		aws.String("ami-0b549f36789a5c4ed"),
		aws.String("ami-0a81acee8067b3f65"),
		aws.String("ami-0fa191704e4fd3aa7"),
		aws.String("ami-05bfef86a955a699e"),
		aws.String("ami-06fe35b1800a64202"),
		aws.String("ami-0fe309367b3f69426"),
		aws.String("ami-0f39a43f9be424601"),
		aws.String("ami-063d90b91a70384a6"),
		aws.String("ami-0939bde2bdaa1f3dd"),
		aws.String("ami-0bc44f20733075c33"),
		aws.String("ami-0b250824c709a5337"),
		aws.String("ami-01fa948bd042d386c"),
		aws.String("ami-04af1eebb10e2cc35"),
		aws.String("ami-06ee372a8602c81c7"),
		aws.String("ami-072f44695fdd32499"),
		aws.String("ami-0491564077c6e3a0a"),
		aws.String("ami-0b2dbefebb844557f"),
		aws.String("ami-0a068d2ab666c55ac"),
		aws.String("ami-01e26c696309a0b1a"),
		aws.String("ami-0108528ba59a156f4"),
		aws.String("ami-06ee5ff817eed44ba"),
		aws.String("ami-06e90f3404ebc4277"),
		aws.String("ami-0d52d2e288d74033c"),
		aws.String("ami-02a0b9d2130102f9c"),
		aws.String("ami-09cc35852d791909d"),
		aws.String("ami-0f228d0f95593ca28"),
		aws.String("ami-038d64845a58ef1e3"),
		aws.String("ami-0ec42c2e3f96bb56a"),
		aws.String("ami-0da962c6867010b60"),
		aws.String("ami-03537d3f9863df131"),
		aws.String("ami-054d1fd63c8670fd9"),
		aws.String("ami-025b1803dd2b90d58"),
		aws.String("ami-082ca0e69db2b7601"),
		aws.String("ami-00bd8add1a2d911df"),
		aws.String("ami-086f2062319bb79e1"),
		aws.String("ami-07a87dd0dd0b306ee"),
	},
	"eu-north-1": {
		aws.String("ami-0f960c8194f5d8df5"),
		aws.String("ami-0cbfcdb45dcced1ca"),
		aws.String("ami-0a6351192ce04d50c"),
		aws.String("ami-064087b8d355e9051"),
		aws.String("ami-0cf13cb849b11b451"),
		aws.String("ami-0232b5b69baa22da3"),
		aws.String("ami-055cbf730c068c80c"),
		aws.String("ami-001adcfb823b62ada"),
		aws.String("ami-01c44fa2d665dfbad"),
		aws.String("ami-0aae2a74b00fb9464"),
		aws.String("ami-0e9469317b5ab5ede"),
		aws.String("ami-095f09fe22186a9f2"),
		aws.String("ami-0673ea1fc7800efdc"),
		aws.String("ami-02c68996dd3d909c1"),
		aws.String("ami-0619bf8a652f4434c"),
		aws.String("ami-0bcb307505fea8239"),
		aws.String("ami-064e4de8fb34b7c17"),
		aws.String("ami-0493d62fb10c2c88a"),
		aws.String("ami-04c3d40468dfd9312"),
		aws.String("ami-0994c54c51db5d17d"),
		aws.String("ami-0c4af0280f45c5a65"),
		aws.String("ami-02f18278b1152528d"),
		aws.String("ami-05682775fab780826"),
		aws.String("ami-0b9b8b6a67b1b959f"),
		aws.String("ami-094ef9f1ca2355632"),
		aws.String("ami-0e15b61fb5a8698b5"),
		aws.String("ami-0c7d85b6a7eb01fda"),
		aws.String("ami-0a90ec058b026a372"),
		aws.String("ami-086eed1a32074a1af"),
		aws.String("ami-00c3f041698a6126b"),
		aws.String("ami-03a2ff446d5bf5187"),
		aws.String("ami-0dffb5995dd0149d8"),
		aws.String("ami-004315db34dc3898d"),
		aws.String("ami-0bf2c695aa8dd07d4"),
		aws.String("ami-0f36f0a5c243b9219"),
		aws.String("ami-0c65f73761a20336a"),
		aws.String("ami-09da0b914d129c078"),
		aws.String("ami-07dfbfd3b7878493f"),
		aws.String("ami-0daa80855adfc10ee"),
		aws.String("ami-05e995d922f2c9a05"),
		aws.String("ami-090b43415c6d3124c"),
		aws.String("ami-0ad0b230154083da2"),
		aws.String("ami-032d1d8ac5df8edfa"),
		aws.String("ami-09ec71e87b5168c02"),
		aws.String("ami-0b907c2c505a4804d"),
		aws.String("ami-063b92755781d63a5"),
		aws.String("ami-0dcd26ce867fb2f9e"),
		aws.String("ami-0e291e6768ab2c00f"),
		aws.String("ami-05712239ec0999b59"),
		aws.String("ami-0c5df76e04e8abe98"),
		aws.String("ami-0ebd4ff61233330ca"),
	},
	"sa-east-1": {
		aws.String("ami-01f451f00dae38302"),
		aws.String("ami-0ce34cc89da9d7a19"),
		aws.String("ami-0e57b6dbf82e26537"),
		aws.String("ami-0b7af114fb404cd23"),
		aws.String("ami-0b3c794a454788ac7"),
		aws.String("ami-0ea67dc1f80ae4aef"),
		aws.String("ami-002a875adefcee7fc"),
		aws.String("ami-07f98243a3ec0cc95"),
		aws.String("ami-0eb55b46af3d99794"),
		aws.String("ami-0867b94bacc5d60ce"),
		aws.String("ami-057c5dec621bf6bca"),
		aws.String("ami-0773230995f88ed15"),
		aws.String("ami-096ea6c54ed7ade4d"),
		aws.String("ami-0d6d5992aee2fb403"),
		aws.String("ami-00598d1087434a359"),
		aws.String("ami-0226d6a40fd39e105"),
		aws.String("ami-063098f580b80b58e"),
		aws.String("ami-028a4fc19cb87274f"),
		aws.String("ami-0a07843c838f0b5ce"),
		aws.String("ami-0bb435fe3bddb1872"),
		aws.String("ami-05610edb12e131601"),
		aws.String("ami-09f5c5e852ea25dd5"),
		aws.String("ami-0e5d58c05b7f1ac68"),
		aws.String("ami-02c6fd63592f81f3c"),
		aws.String("ami-0d304756c4714467b"),
		aws.String("ami-0fd58b2e7a5e738cc"),
		aws.String("ami-0688b3ba053f7eb5b"),
		aws.String("ami-0425296d62cbb9439"),
		aws.String("ami-0dca2f35499b986f0"),
		aws.String("ami-0d034087dde611982"),
		aws.String("ami-064fa7ea182ab724c"),
		aws.String("ami-0b87061262a15346a"),
		aws.String("ami-0eb55b46af3d99794"),
		aws.String("ami-0b123273fd25fe833"),
		aws.String("ami-063641621986c61d4"),
		aws.String("ami-05db9c2355aa38f0d"),
		aws.String("ami-05ac44b8b463d3d19"),
		aws.String("ami-0136df87b8e466067"),
		aws.String("ami-05486c997447b34b4"),
		aws.String("ami-0eef8c7aff5baa803"),
		aws.String("ami-09ee39f1f0f164892"),
		aws.String("ami-068e2d4bc28a6ba23"),
		aws.String("ami-03315642ca03e6ab8"),
		aws.String("ami-0fa8efd4f56c668b6"),
		aws.String("ami-0c00453583aaf434e"),
		aws.String("ami-0bf606f6236128bd0"),
	},
	"us-east-1": {
		aws.String("ami-06e46074ae430fba6"),
		aws.String("ami-0fa1de1d60de6a97e"),
		aws.String("ami-016eb5d644c333ccb"),
		aws.String("ami-02978b79564e08f2f"),
		aws.String("ami-007855ac798b5175e"),
		aws.String("ami-0aa2b7722dc1b5612"),
		aws.String("ami-068f27965379d536b"),
		aws.String("ami-0e38fa17744b2f6a5"),
		aws.String("ami-03a21b62905737826"),
		aws.String("ami-0b7dd7b9e977b2b85"),
		aws.String("ami-0e9e596470c5c4caa"),
		aws.String("ami-0661ef50fa4f6c748"),
		aws.String("ami-00b2c40b15619f518"),
		aws.String("ami-06b16fbd550e52604"),
		aws.String("ami-00fdc7a8593bd6e9d"),
		aws.String("ami-0dacd9d37849a9f38"),
		aws.String("ami-09eee5f3b0e4caa8b"),
		aws.String("ami-0d5a3477c63b82c82"),
		aws.String("ami-08ebc8244f9fa2216"),
		aws.String("ami-0e82ebb170891987a"),
		aws.String("ami-04c53928a475d72c2"),
		aws.String("ami-0b31c97165757d1a2"),
		aws.String("ami-0db4bed99e0cb80b7"),
		aws.String("ami-0e0a031112af8987e"),
		aws.String("ami-0b5a9ac7071701572"),
		aws.String("ami-070cd510a2f8432fc"),
		aws.String("ami-0b87e73388f8a4fd0"),
		aws.String("ami-0e58f89e91723af4c"),
		aws.String("ami-005b11f8b84489615"),
		aws.String("ami-0874d82d2138e9fd1"),
		aws.String("ami-0fc522222ab74a244"),
		aws.String("ami-0989629f01a1cb33e"),
		aws.String("ami-0601b11fc97febe3c"),
		aws.String("ami-085a3abb84068d568"),
		aws.String("ami-0d1ecaad2613d97c3"),
		aws.String("ami-007cf291af489ad4d"),
		aws.String("ami-0c6c29c5125214c77"),
		aws.String("ami-004811053d831c2c2"),
	},
	"us-east-2": {
		aws.String("ami-0103f211a154d64a6"),
		aws.String("ami-0d80c4e4338722fc6"),
		aws.String("ami-067a8829f9ae24c1c"),
		aws.String("ami-0a04068a95e6a1cde"),
		aws.String("ami-0fb3a91b7ce257ec1"),
		aws.String("ami-0a695f0d95cefc163"),
		aws.String("ami-0ca58e4cb9e83244e"),
		aws.String("ami-0a090038a7c772680"),
		aws.String("ami-0b5f6c4fef0c632bc"),
		aws.String("ami-0326a3ddb7d831ad5"),
		aws.String("ami-0d8aa9577a621441a"),
		aws.String("ami-0f35413f664528e13"),
		aws.String("ami-0253800445e4ada85"),
		aws.String("ami-05df72f472ebf0066"),
		aws.String("ami-00d7773c227e08595"),
		aws.String("ami-0bb258892aeaac588"),
		aws.String("ami-042f6473eb1e6af71"),
		aws.String("ami-0d56c020aa1219f11"),
		aws.String("ami-0c4f2b99b40a37af8"),
		aws.String("ami-071f5a56c6426f504"),
		aws.String("ami-00643bb24a9b530b0"),
		aws.String("ami-050315e917a471a11"),
		aws.String("ami-07d620bb2a33317bf"),
		aws.String("ami-065a34b2891054895"),
		aws.String("ami-05f6084a6b524a8f0"),
		aws.String("ami-06c4532923d4ba1ec"),
		aws.String("ami-0fddd1d73e7a86674"),
		aws.String("ami-015c77c4b4a67541e"),
		aws.String("ami-0bea0be466a642cf0"),
		aws.String("ami-01965da38968348f8"),
		aws.String("ami-05eea9de06c839c7c"),
		aws.String("ami-08778753ef37aa408"),
		aws.String("ami-0a4f85fc401a53c3d"),
		aws.String("ami-068ee501b16366120"),
		aws.String("ami-0c77e0bf18f26b6db"),
		aws.String("ami-080935679c2c240a0"),
		aws.String("ami-07d16074c2fdf3a19"),
		aws.String("ami-0b54bb4e237a21bfc"),
		aws.String("ami-06ad5df4b191f59d0"),
		aws.String("ami-052fd3067d337faf6"),
		aws.String("ami-0af198159897e7a29"),
		aws.String("ami-065eed752f466bb1c"),
	},
	"us-west-1": {
		aws.String("ami-09c5c62bac0d0634e"),
		aws.String("ami-0e5e4bbcbd7349cac"),
		aws.String("ami-0e16c3bc75f23e32b"),
		aws.String("ami-00522a6964c4a0c59"),
		aws.String("ami-014d05e6b24240371"),
		aws.String("ami-0dc50976db563e50a"),
		aws.String("ami-081a3b9eded47f0f3"),
		aws.String("ami-01d9361b1c190a61b"),
		aws.String("ami-0112ad2d371ed0d9a"),
		aws.String("ami-06393c0d07516eeb0"),
		aws.String("ami-0d6776f702f95c08c"),
		aws.String("ami-09ffef517766fd3da"),
		aws.String("ami-0bf166b48bbe2bf7c"),
		aws.String("ami-0d39d92a8fc317f94"),
		aws.String("ami-0e62435e1c26888b7"),
		aws.String("ami-08bb2e1d9525135ab"),
		aws.String("ami-0366f637b0605ddea"),
		aws.String("ami-0d4d9184487a7c59c"),
		aws.String("ami-0251e72a026ed1705"),
		aws.String("ami-05d82cff82fa19e79"),
		aws.String("ami-00088539b71c6191c"),
		aws.String("ami-06e15c3858febd415"),
		aws.String("ami-05ef351f130b40597"),
		aws.String("ami-0e056dea26af2752a"),
		aws.String("ami-099d5c61edbb230ca"),
		aws.String("ami-038c096a4a922054d"),
		aws.String("ami-07f83aa8dc0abea00"),
		aws.String("ami-025897efac0d2f104"),
		aws.String("ami-0c43eb5d8849af79f"),
		aws.String("ami-06393c0d07516eeb0"),
		aws.String("ami-0267fc24ee0102728"),
		aws.String("ami-036653a4cdb4c714b"),
		aws.String("ami-0d44f1e45c689ccf7"),
		aws.String("ami-0e22afe60ccbcd9e0"),
		aws.String("ami-040171276b21e6ec5"),
		aws.String("ami-0c3ce577b7085056d"),
		aws.String("ami-06a54aa5e58a60898"),
		aws.String("ami-0b7719376e93ddfd7"),
		aws.String("ami-042b6a9e11afe8322"),
		aws.String("ami-06ce593d0cc7740d7"),
		aws.String("ami-0bd3b255f1beeae5e"),
		aws.String("ami-0c75abbbe8bd81092"),
		aws.String("ami-0d05dc16022b76cb6"),
	},
	"us-west-2": {
		aws.String("ami-0747e613a2a1ff483"),
		aws.String("ami-0c252bb9e6b71848e"),
		aws.String("ami-0dda7e535b65b6469"),
		aws.String("ami-019866531dac971b9"),
		aws.String("ami-079ba66a5e9f2b70e"),
		aws.String("ami-0fcf52bcf5db7b003"),
		aws.String("ami-0db245b76e5c21ca1"),
		aws.String("ami-0b8508d946afbf024"),
		aws.String("ami-04861a2e4ef4abd6b"),
		aws.String("ami-0e8aab9a7a911cdc0"),
		aws.String("ami-0a39ed2b865d65970"),
		aws.String("ami-0de8d936951b16335"),
		aws.String("ami-06b8d5099f3a8d79d"),
		aws.String("ami-0c1b4dff690b5d229"),
		aws.String("ami-0f447c5d18f9110f9"),
		aws.String("ami-00719b15124c74012"),
		aws.String("ami-08824607caa2b141e"),
		aws.String("ami-0e4d5442b2f304425"),
		aws.String("ami-0a379fdc08c58a556"),
		aws.String("ami-035c1d9e7a1fca683"),
		aws.String("ami-032abe919d133fa58"),
		aws.String("ami-00626ada821398eef"),
		aws.String("ami-06132d0557a481472"),
		aws.String("ami-0de611416bfdc2235"),
		aws.String("ami-0718598d42c919a07"),
		aws.String("ami-03fa57f28962c67c1"),
		aws.String("ami-0a242608b33937c51"),
		aws.String("ami-0521250034feb12a6"),
		aws.String("ami-07e5f0693976c0c2e"),
		aws.String("ami-09a728481ab9015b5"),
		aws.String("ami-033f3dfa0c1c584a4"),
		aws.String("ami-081aaface2871d0d0"),
		aws.String("ami-03ff9501b439808eb"),
		aws.String("ami-012751beb242b0262"),
		aws.String("ami-0a75b40c9f7cd0a2a"),
		aws.String("ami-0b24eb6af9b3438b4"),
		aws.String("ami-05f76944fb02a96c6"),
		aws.String("ami-049602f1e5c535667"),
		aws.String("ami-0a4237d2123224da7"),
		aws.String("ami-024148801f88be4a4"),
		aws.String("ami-014597443ebdf8500"),
		aws.String("ami-05e34ea549fddf826"),
		aws.String("ami-08911268ee09cb08e"),
		aws.String("ami-00f902c807805f51a"),
		aws.String("ami-03f6bd8c9c6230968"),
		aws.String("ami-0a7511ae140b1aa4d"),
		aws.String("ami-0572bc6dc45296d87"),
	},
	"ap-east-1": {
		aws.String("ami-00292313920481de3"),
		aws.String("ami-0e5ee44e38fd711f7"),
		aws.String("ami-08231fa88fd6ad25c"),
		aws.String("ami-0f3d30a05b4b0541b"),
		aws.String("ami-0b15c99a3289fb1dc"),
		aws.String("ami-0c5380960dac0cf50"),
		aws.String("ami-0d7ce860e738db09b"),
		aws.String("ami-02d83de59e3879a4f"),
		aws.String("ami-0331581d6dc086008"),
		aws.String("ami-0691bf13b7b64b8cd"),
		aws.String("ami-087fedc69aa2a97b2"),
		aws.String("ami-03cd9ca5c6dec4980"),
		aws.String("ami-055fe8a1e891fd63b"),
		aws.String("ami-0a0a1faa79405fea1"),
		aws.String("ami-01384c3926bc6f380"),
		aws.String("ami-0da1beb0e7263e7e1"),
		aws.String("ami-0ee3739f93544eb5e"),
		aws.String("ami-0d31804aad3f9cac9"),
		aws.String("ami-09743707599051b74"),
		aws.String("ami-075534f51749adafd"),
		aws.String("ami-089145b36400d5de9"),
		aws.String("ami-0382f5a38541291e9"),
		aws.String("ami-084994542408d6915"),
		aws.String("ami-096453646d88249b7"),
		aws.String("ami-0a1248aee5972605c"),
		aws.String("ami-0f4091e2829cd531c"),
		aws.String("ami-0b4f8444c1d0335c3"),
		aws.String("ami-0400b37e022650e20"),
		aws.String("ami-000e52b31c4b4caa5"),
		aws.String("ami-0b34526e2e976a254"),
		aws.String("ami-02aefa095e7ff4735"),
		aws.String("ami-07dfd30ae10303c2b"),
		aws.String("ami-00145bfaf93729d9d"),
		aws.String("ami-01ebc8e4ee32b7f59"),
		aws.String("ami-0001eb198f16f76cc"),
		aws.String("ami-00b5070132a77fb42"),
		aws.String("ami-0126a410bd65df493"),
		aws.String("ami-07c2714ba0b46274e"),
		aws.String("ami-0ff46522157342dd0"),
		aws.String("ami-0e12288321380f14e"),
		aws.String("ami-0d22c95144be9606a"),
		aws.String("ami-0512a9e52dc0b073f"),
		aws.String("ami-0e5a707797944044b"),
	},
	"ap-south-1": {
		aws.String("ami-07d3a50bd29811cd1"),
		aws.String("ami-09f7fbc41963e146f"),
		aws.String("ami-0fdea1353c525c182"),
		aws.String("ami-0466079baeb1010da"),
		aws.String("ami-0ce4c5d21b8e0de83"),
		aws.String("ami-02eb7a4783e7e9317"),
		aws.String("ami-09461328af8fbcb9c"),
		aws.String("ami-071bfb668966c44ae"),
		aws.String("ami-0b156a71992057092"),
		aws.String("ami-03a933af70fa97ad2"),
		aws.String("ami-0839f419042b6d609"),
		aws.String("ami-09a32118945ae6b71"),
		aws.String("ami-0f8eab2ed2d4eeebb"),
		aws.String("ami-04923cf973ad5a828"),
		aws.String("ami-0695b18c741b3999f"),
		aws.String("ami-0e50b9b105ec75754"),
		aws.String("ami-0746170310ebe2359"),
		aws.String("ami-0193eebd3608becde"),
		aws.String("ami-0cd4a94579fa4cb4f"),
		aws.String("ami-0066504222a9077b9"),
		aws.String("ami-05f83ffa8a62c0009"),
		aws.String("ami-047142359f216747f"),
		aws.String("ami-01b5038bb2f61e6a0"),
		aws.String("ami-09e11ca44251a75b9"),
		aws.String("ami-05ee94740c98ea632"),
		aws.String("ami-0fafaf9a99d9e9082"),
		aws.String("ami-0ccbf64927ca7d108"),
		aws.String("ami-07e672ad222f2f48d"),
		aws.String("ami-004b958968a74efe8"),
		aws.String("ami-0ef5a37e767842839"),
		aws.String("ami-061183ad486d5dd8a"),
		aws.String("ami-014d846d6e42c37b5"),
		aws.String("ami-099770dc1e84b8125"),
		aws.String("ami-08afb9d55a5dbd657"),
		aws.String("ami-07a89acc76f0335e1"),
		aws.String("ami-00061bb9f9cd71566"),
		aws.String("ami-04daff085607f4847"),
		aws.String("ami-058ad16b95eb47717"),
		aws.String("ami-0d38850c2ecc8a352"),
		aws.String("ami-0754052f3e64ffb65"),
		aws.String("ami-0a5dcff6fb7af3fc9"),
		aws.String("ami-0dddd4d17b84ac14c"),
		aws.String("ami-06ac5f5ed93f43a1d"),
	},
	"ap-northeast-3": {
		aws.String("ami-04f309177872e4a43"),
		aws.String("ami-0f19932fc7b89784a"),
		aws.String("ami-085d2183d3c2e3576"),
		aws.String("ami-0134aa607f004fae1"),
		aws.String("ami-086badb75ca1dd7c8"),
		aws.String("ami-083b39ad40b899755"),
		aws.String("ami-07a129a553165490c"),
		aws.String("ami-07dd60ade3817c079"),
		aws.String("ami-05dc534011a0c80f5"),
		aws.String("ami-0f9015c078d524d7a"),
		aws.String("ami-0ab6ea481909397f7"),
		aws.String("ami-07c300a9f57f8f5cb"),
		aws.String("ami-0b80896a0bdb31214"),
		aws.String("ami-0ad5ac148612da32d"),
		aws.String("ami-0f5ac783bf46945a1"),
		aws.String("ami-0686de9e11ae8547c"),
		aws.String("ami-0cef2102966e16342"),
		aws.String("ami-0bbc7eecbde53a6a6"),
		aws.String("ami-0ff772fa309b750ac"),
		aws.String("ami-081bd0b1d7c6ef2a9"),
		aws.String("ami-08bbadbb139bb7547"),
		aws.String("ami-0bb203334132934cf"),
		aws.String("ami-01a41e9269558edfe"),
		aws.String("ami-01c33481c181485be"),
		aws.String("ami-04055196110a31ee7"),
		aws.String("ami-0ccc6c2fae7e6d3f4"),
		aws.String("ami-0447a13fa007e018b"),
		aws.String("ami-01d7109b399c6b223"),
		aws.String("ami-067312d8a1eecc82b"),
		aws.String("ami-092949e55d24ec6a6"),
		aws.String("ami-005cdeee75db4b9f2"),
		aws.String("ami-0f50e02fb405b71f2"),
		aws.String("ami-0958b319e20935da2"),
		aws.String("ami-0712cee649ce32f95"),
		aws.String("ami-0b5ae7c57d9a8238f"),
		aws.String("ami-0c721a3794ea9c88b"),
		aws.String("ami-098ecf0f7af6eccaf"),
		aws.String("ami-09c0e9092a7cc36b1"),
		aws.String("ami-0070738bace68ab78"),
		aws.String("ami-0bc1ccedf8088196f"),
		aws.String("ami-0f0f7b2d5100f1167"),
		aws.String("ami-01c9f367ff89c5cac"),
		aws.String("ami-0956b03b0cacb9fdc"),
	},
	"ap-northeast-2": {
		aws.String("ami-0676d41f079015f32"),
		aws.String("ami-0e52aed83baf3e36a"),
		aws.String("ami-0f5b323525724e5ca"),
		aws.String("ami-0e8624bf9e85611f8"),
		aws.String("ami-089d769d0a18be35b"),
		aws.String("ami-0b1f7127e83f3de59"),
		aws.String("ami-04cebc8d6c4f297a3"),
		aws.String("ami-0c6e5afdd23291f73"),
		aws.String("ami-04d79297c23b8d80c"),
		aws.String("ami-0b6f6543c11f4ab41"),
		aws.String("ami-06940e714201cabec"),
		aws.String("ami-0dc6794619a213e75"),
		aws.String("ami-012fdadbf1573c752"),
		aws.String("ami-0e77b0d1fca185fa9"),
		aws.String("ami-08d910826a226eaf7"),
		aws.String("ami-055437ee4dcf92427"),
		aws.String("ami-015b827566ddeb79f"),
		aws.String("ami-0061803186499c438"),
		aws.String("ami-0dbd4dcb554e2deaf"),
		aws.String("ami-0cbba000fd832fcbf"),
		aws.String("ami-018e7b32352327666"),
		aws.String("ami-0854534cddfc9644f"),
		aws.String("ami-064a230b4f4b64b45"),
		aws.String("ami-0d7941e2df516d44b"),
		aws.String("ami-01ca18b2668eb76f7"),
		aws.String("ami-0ccce9ef7dd463bff"),
		aws.String("ami-06e46074ae430fba6"),
		aws.String("ami-085a3abb84068d568"),
		aws.String("ami-0f1bf0c7b0f2c8d3e"),
		aws.String("ami-0a0c8d7f5d1a2739c"),
		aws.String("ami-06b9122710049dfe7"),
		aws.String("ami-009a7a27f88698977"),
		aws.String("ami-0717542e143fd7d10"),
		aws.String("ami-08a60fe0878be499d"),
		aws.String("ami-03dbf0ed1afdb7d85"),
		aws.String("ami-0bc2548e307664f20"),
		aws.String("ami-0e8f8aaacc1b5b785"),
		aws.String("ami-0fbe7c605672deec0"),
		aws.String("ami-081aee8bd16834c03"),
		aws.String("ami-064fb983dc18d5f67"),
		aws.String("ami-085e05b0ebc9565d4"),
		aws.String("ami-08af8fcc742b573b3"),
		aws.String("ami-0ef2869cfb6341803"),
		aws.String("ami-062c81e3fe13b4a06"),
		aws.String("ami-0ede40e86aa32183a"),
		aws.String("ami-0084ff4520f7fd92a"),
		aws.String("ami-057de62f4b76a6c17"),
		aws.String("ami-039afd24b82f2b702"),
		aws.String("ami-0fbe49b81e628ec6b"),
	},
	"ap-southeast-1": {
		aws.String("ami-04ddf30efb5385f93"),
		aws.String("ami-02b2e78e9b867ffec"),
		aws.String("ami-04ba270ccd8098407"),
		aws.String("ami-08510a2cd06f96e9a"),
		aws.String("ami-000e680c06e6ad00b"),
		aws.String("ami-0a72af05d27b49ccb"),
		aws.String("ami-062550af7b9fa7d05"),
		aws.String("ami-0e42bfd2029a917a4"),
		aws.String("ami-0370427aca55b362e"),
		aws.String("ami-0af87cf8bd1471a46"),
		aws.String("ami-0e2056101d78bef54"),
		aws.String("ami-0eb00265474e95a37"),
		aws.String("ami-0010144f2cd5aed03"),
		aws.String("ami-058ba3653751bf529"),
		aws.String("ami-0fcd2e9ac9a168217"),
		aws.String("ami-02eda3a049e7c3cc8"),
		aws.String("ami-0d9d1c1d8e2c986a4"),
		aws.String("ami-072f48a9ed4f1bbda"),
		aws.String("ami-0c6ebf4e6fe527243"),
		aws.String("ami-03be746827dd5f119"),
		aws.String("ami-09c209043e0252301"),
		aws.String("ami-04bc6caca307e18da"),
		aws.String("ami-045d2f7296c1129df"),
		aws.String("ami-0c7486b2f37da2312"),
		aws.String("ami-01ae4e0d66e6d1b46"),
		aws.String("ami-088235c4ac421ec9d"),
		aws.String("ami-0446a4129196f932e"),
		aws.String("ami-044ed7e0c9994383e"),
		aws.String("ami-0eae6d9e939614228"),
		aws.String("ami-0ac7ca39f7d56f7c3"),
		aws.String("ami-0f084631dbce400ed"),
		aws.String("ami-054a35787034134d2"),
		aws.String("ami-066938299d0f4c4a0"),
		aws.String("ami-0bbc0fa1a762e63c9"),
		aws.String("ami-0ac6514823583e0f1"),
		aws.String("ami-0f7199970f5bdc22c"),
		aws.String("ami-0d64fe8ab57f2eaea"),
		aws.String("ami-0f393f5796dd4ade3"),
		aws.String("ami-0634a45793db5c06b"),
		aws.String("ami-0dd28c803f9318319"),
		aws.String("ami-0164ce74ec6eb54e9"),
		aws.String("ami-0076ac6f47beccc0d"),
	},
}

// ListImage ...
// reference: https://docs.amazonaws.cn/AWSEC2/latest/APIReference/API_DescribeImages.html
func (a *Aws) ListImage(kt *kit.Kit, opt *image.AwsImageListOption) (*image.AwsImageListResult, error) {
	client, err := a.clientSet.ec2Client(opt.Region)
	if err != nil {
		return nil, err
	}

	imageIDs, exist := regionImageIDMap[opt.Region]
	if !exist {
		return &image.AwsImageListResult{
			Details:   make([]image.AwsImage, 0),
			NextToken: nil,
		}, nil
	}

	req := &ec2.DescribeImagesInput{MaxResults: opt.Page.MaxResults, NextToken: opt.Page.NextToken}
	req.Filters = []*ec2.Filter{
		{Name: aws.String("is-public"), Values: []*string{aws.String("true")}},
		{Name: aws.String("state"), Values: []*string{aws.String("available")}},
		{
			Name:   aws.String("image-id"),
			Values: imageIDs,
		},
	}

	resp, err := client.DescribeImagesWithContext(kt.Ctx, req)
	if err != nil {
		logs.Errorf("describe aws image failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	images := make([]image.AwsImage, 0)
	for _, pImage := range resp.Images {
		images = append(images, image.AwsImage{
			CloudID:      *pImage.ImageId,
			Name:         *pImage.Name,
			State:        *pImage.State,
			Architecture: *pImage.Architecture,
			Type:         "public",
		})
	}
	return &image.AwsImageListResult{Details: images, NextToken: resp.NextToken}, nil
}
