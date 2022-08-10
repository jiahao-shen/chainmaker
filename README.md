# ChainMaker

## 长安链搭建
1. 进入工作目录: `cd ${WORKDIR}`
2. 下载项目: `git clone -b v2.2.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-go.git`
3. 下载项目: `git clone -b v2.2.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-cryptogen.git`
4. 编译证书: `cd chainmaker-cryptogen && make`
5. 软链接: `ln -s ${WORKDIR}/chainmaker-cryptogen ${WORKDIR}/chainmaker-go/tools`
6. 生成单链单节点证书配置文件:
   - `cd ${WORKDIR}/chainmaker-go/scripts`
   - `./prepare.sh 1 1`
     - `input consensus type (0-SOLO,1-TBFT(default),3-MAXBFT,4-RAFT): 0`
     - `input log level (DEBUG|INFO(default)|WARN|ERROR): INFO`
     - `enable docker vm (YES|NO(default))YES`
   - `tree -L 3 ../build`
7. 编译安装包:
   - `./build_release.sh`
   - `tree ../build/release`
8. 解压安装包并运行:
   - `cd ../build/release`
   - `tar zxvf chainmaker-xxxx.tar.gz`
   - `docker pull chainmakerofficial/chainmaker-vm-docker-go:v2.2.0.1`
   - `cd chainmaker-xxxx/bin && ./start.sh`

<!-- 8. 运行节点集群:
   - 启动: `./cluster_quick_start.sh normal`
   - 查看进程: `ps -ef|grep chainmaker | grep -v grep`
   - 查看端口:
     - (Linux)`netstat -lptn | grep 1230`
     - (macOS)`netstat -an | grep 1230` -->

## CMC命令行工具安装
1. `cd ${WORKDIR}/chainmaker-go/tools/cmc`
2. `go build`

### 合约创建
```
./cmc client contract user create \
--contract-name=contract_fact \
--runtime-type=DOCKER_GO \
--byte-code-path=./testdata/docker-go-demo/contract_fact.7z \
--version=1.0 \
--sdk-conf-path=./testdata/sdk_config.yml \
--admin-key-file-paths=./testdata/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.key,./testdata/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.key,./testdata/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.key,./testdata/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.key \
--admin-crt-file-paths=./testdata/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.crt,./testdata/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.crt,./testdata/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.crt,./testdata/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.crt \
--sync-result=true \
--params="{}"
```
### 调用合约
```
./cmc client contract user invoke \
--contract-name=contract_fact \
--method=invoke_contract \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"method\":\"save\",\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" \
--sync-result=true
```

### 查询合约
```
./cmc client contract user get \
--contract-name=contract_fact \
--method=invoke_contract \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"method\":\"findByFileHash\",\"file_hash\":\"ab3456df5799b87c77e7f88\"}"
```

## Docker-Go合约开发
1. 在目录`chainmaker-contract-sdk-docker-go`中编写合约
2. 执行`buid.sh`编译合约, 生成的文件存放于`target`目录下

## Go-SDK使用教程
1. 编写配置文件, 样例见`sdk-go/client/sdk_config.yaml`
2. 确认证书`crypto-config`的路径
3. 确认编译后的7z合约路径,
4. 安装, 查询, 调用合约方法见`sdk-go/client/sdk_test.go`文件中的示例
5. !!!注意: `QueryContract`和`InvokeContract`中的method参数需设置为"invoke_contract", 具体调用的方法包装在params字段中
