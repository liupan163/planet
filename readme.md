# planet

**planet** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see
the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install
dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see
the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the
configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/planet@latest! | sudo bash
```

`username/planet` should match the `username` and `repo_name` of the Github repository to which the source code was
pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## cmd

- ignite scaffold chain planet --no-module
- ignite scaffold module blog --ibc  添加blog，并且ibc module_ibc.go
- ignite scaffold list post title content creator --no-message --module blog
- ignite scaffold list sentPost postID title chain creator --no-message --module blog
- ignite scaffold list timeoutPost title chain creator --no-message --module blog

- ignite scaffold packet ibcPost title content --ack postID --module blog
  - 定义了一个数据结构ibcPost，发送包括：title、content，返回包括：postID。

proto/planet/blog/packet.proto
增加  string creator = 2;
ignite chain build


ignite chain serve -c earth.yaml
ignite chain serve -c mars.yaml

rm -rf ~/.ignite/relayer

ignite relayer connect

## 
 
课后作业
1.参考该文档https://docs.ignite.com/guide/install， 在自己的电脑上安装Ignite CLI环境。

2.按照课堂上演示的案例，重复每条步骤，在本地生成一个同样的项目。（在生成的过程中，确保理解每个步骤的意义，通过生成的代码变化，弄清楚每条命令具体做了什么）

3.创建一个新的ibc数据包，名称为updatePost，包含postID、title、content。
通过将该数据包发送给对手链，可更新存储在对手链中某条Post（通过postID）的标题和内容。
对手链在ack确认数据包中将返回是否更新成功，发送链收到确认数据包后更新本地之前存储的SentPost为新的title。

- ignite scaffold packet ibcPost title content --ack postID --module blog
- ignite scaffold packet updatePost postID title content --ack isUpdateOk --module blog


4.启动中继器和两条链节点，发送一条post给对手链，发送成功后，确保对手链能查询到该post以及当前链保存了相应的sentPost。获取该post对应的postID，向对手链发送一条updatePost的ibc数据包，更新该post的title和content。发送成功后，查询对手链的post和当前链的sentPost，确保更新成功。


一点提示：
1.若根据postID无法找到对应的post，则返回更新失败。
2.对于post和sentPost的查询和更新，在第二步中都有相关代码生成，不需要额外开发。
3.修改完成后需要重新编译planetd。