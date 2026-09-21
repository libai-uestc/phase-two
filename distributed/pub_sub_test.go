package distributed_test

import (
	"context"
	"libai/go/phase-two/distributed"
	"testing"
)

func TestPubSub(t *testing.T) {
	distributed.PubSub(context.Background(), client)
}

// go test -v .\distributed\ -run=^TestPubSub$ -count=1
/*

2026/09/21 13:56:16 INFO connect to redis
=== RUN   TestPubSub
%!s(<nil>)从频道channel1里接收到消息:奔流到海不复回
%!s(<nil>)向频道channel1发布了消息，此时频道也1个订阅者
%!s(<nil>)从频道channel1里接收到消息:黄河之水天上来
%!s(<nil>)向频道channel1发布了消息，此时频道也1个订阅者
--------------------------------------------------
%!s(<nil>)从频道channel2里接收到消息:千金散尽还复来
%!s(<nil>)从频道channel2里接收到消息:天生我材必有用
%!s(<nil>)向频道channel2发布了消息，此时频道也2个订阅者
%!s(<nil>)向频道channel2发布了消息，此时频道也2个订阅者
%!s(<nil>)从频道channel2里接收到消息:千金散尽还复来
%!s(<nil>)从频道channel2里接收到消息:天生我材必有用
--- PASS: TestPubSub (4.00s)
PASS
ok      libai/go/phase-two/distributed  5.396s
PS C:\Users\18101\Desktop\第二阶段\phase-two>
*/
