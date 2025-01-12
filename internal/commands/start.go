package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
)

func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	dispatcher.AddHandler(handlers.NewCommand("start", start))
}

func start(ctx *ext.Context, u *ext.Update) error {
	chatId := u.EffectiveChat().GetID()
	peerChatId := ctx.PeerStorage.GetPeerById(chatId)
	if peerChatId.Type != int(storage.TypeUser) {
		return dispatcher.EndGroups
	}
	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
		ctx.Reply(u, "白名单模式已开启 你不在白名单内 请找 @listder 申请", nil)
		return dispatcher.EndGroups
	}
	if len(config.ValueOf.BlockUsers) != 0 && utils.Contains(config.ValueOf.BlockUsers, chatId) {
		ctx.Reply(u, "你因为某些原因被封禁了 请找 @listder 解封", nil)
		return dispatcher.EndGroups
	}
	ctx.Reply(u, "把文件发送给我就可以获得下载链接 注意：真人色情内容可能会被封禁", nil)
	return dispatcher.EndGroups
}
