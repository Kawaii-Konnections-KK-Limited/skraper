package skraper_test

import (
	"fmt"
	"testing"

	"github.com/Kawaii-Konnections-KK-Limited/skraper/pkg/skraper"

	"golang.org/x/exp/slices"
)

func TestFindLinksInText(t *testing.T) {
	t.Run("findtext", func(t *testing.T) {
		links := skraper.FindLinksInText(`کانفیگ اهدایی:
		@iran_ray
		
		
		لوکیشن هلند برای دوستانی که درخواست کرده بودند
		
		
		محدودیت حجم: 1 ترابایت
		محدودیت روز:  ندارد
		test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇩🇪 آلمان

🔗 Link : vless://6e326c62-70a1-4f9e-8356-1a9ff6f35ceb@XsV2ray.xalixv2ray.space:443?mode=multi&security=reality&encryption=none&pbk=hbmwoSOK2SM0C_FmhWJJZM-up8d4dDOMscmyCA8gynk&fp=firefox&spx=%2F&type=grpc&serviceName=@XsV2ray,@XsV2ray&sni=www.speedtest.net&sid=7e4e8776#🇩🇪 آلمان : @KralVPN

⚙ isp : Hetzner Online GmbH
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇬🇧 انگلیس

🔗 Link : vless://814ea288-7c9a-4830-b2b1-7fd217cde463@npl9awgerdexsnx.dedicatedlink.co:29039?security=none&encryption=none&headerType=none&type=tcp#🇬🇧 انگلیس : @KralVPN

⚙ isp : Amazon Technologies Inc.
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇮🇷 ایران به خارج (فوروارد ترافیک)

🔗 Link : vless://c9166a33-0a41-4d07-ac61-be20b5e51c2f@online.4zkaban.xyz:80?encryption=none&headerType=none&type=ws&path=%2F%40Ok_Config&host=16.ok-config-31.buzz#🇮🇷 ایران به خارج (فوروارد ترافیک) : @KralVPN

⚙ isp : ANYCAST
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇳🇱 هلند

🔗 Link : vless://dc3e4b38-292b-482c-aa8f-c7f83f88fec3@89.23.103.13:8443?security=reality&encryption=none&pbk=pzMD7WyB7KTYR3UpdevG4lw5WzkJNqt6BzAPJtsNZiU&headerType=none&fp=firefox&spx=%2F&type=grpc&serviceName=Telegram#🇳🇱 هلند : @KralVPN

⚙ isp : Global Internet Solutions LLC
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇨🇦 کانادا

🔗 Link : vless://1e980f99-fa21-4bf9-8b42-496877c944e4@mci.ArV2ray.host:2087?type=grpc&serviceName=@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray&security=tls&fp=&alpn=h2%2Chttp%2F1.1&allowInsecure=1&sni=ArV2ray.Rvin.tech#🇨🇦 کانادا : @KralVPN

⚙ isp : Cloudflare, Inc.
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇷🇺 روسیه

🔗 Link : vless://9fac0e66-13a0-443f-fd7a-8c0b027e4630@mtn.MsV2ray.space:2053?mode=gun&security=tls&encryption=none&alpn=h2,http/1.1&type=grpc&serviceName=@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray,@MsV2ray&sni=tm.MsV2ray.cfd#🇷🇺 روسیه : @KralVPN

⚙ isp : MFC Zaymer LLC
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vmess
📍 لوکیشن : 🇳🇱 هلند

🔗 Link : vmess://eydhZGQnOiAndi1ubDEubmd2aXAubmV0JywgJ2FpZCc6ICcwJywgJ2FscG4nOiAnJywgJ2ZwJzogJycsICdob3N0JzogJycsICdpZCc6ICdkN2FjNDUwYi1hODU4LTQ5M2ItOGM3Ny1kNmFiYjU4ZTkxMTMnLCAnbmV0JzogJ3dzJywgJ3BhdGgnOiAnLycsICdwb3J0JzogJzI2ODMwJywgJ3BzJzogJ/Cfh7Pwn4exINmH2YTZhtivIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJycsICd0bHMnOiAndGxzJywgJ3R5cGUnOiAnJywgJ3YnOiAnMid9

⚙ isp : NForce Entertainment B.V.
📍 @KralVPN

test channel, [8/26/2023 2:38 PM]
💯 کانفیگ Vless
📍 لوکیشن : 🇫🇮 فنلاند

🔗 Link : vless://Telegram--vpnepic@65.109.178.214:443?mode=gun&security=reality&encryption=none&pbk=XY5HUCQudXlSQm4ahgKGqHm8KRH1w0U3VFVUnHmMJB0&fp=chrome&spx=%2F&type=grpc&serviceName=&sni=greenpepper.ir&sid=4e7c606c#🇫🇮 فنلاند : @KralVPN

⚙ isp : Hetzner Online GmbH
📍 @KralVPN
		vmess://eydhZGQnOiAnZWR1Lm1yc3BvdC50aycsICdhaWQnOiAnMCcsICdhbHBuJzogJ2gyLGh0dHAvMS4xJywgJ2ZwJzogJ3JhbmRvbWl6ZWQnLCAnaG9zdCc6ICdlZHUubXJzcG90LnRrJywgJ2lkJzogJzZhYzdkM2VlLTgyMzEtNDM2Ny1hZTFhLTI4ZWNlZDQzY2ZiMScsICduZXQnOiAndGNwJywgJ3BhdGgnOiAnL3RjcDE5MzI2JywgJ3BvcnQnOiAnMTkzMjYnLCAncHMnOiAn8J+HrvCfh7cg2KfbjNix2KfZhiDYqNmHINiu2KfYsdisICjZgdmI2LHZiNin2LHYryDYqtix2KfZgduM2qkpIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJ2VkdS5tcnNwb3QudGsnLCAndGxzJzogJ3RscycsICd0eXBlJzogJ2h0dHAnLCAndic6ICcyJ30=
		vless://1cb85996-a546-4b7a-bc5a-fefc35da7939@mkh.arv2ray.host:2083?security=tls&sni=gr.ArV2ray.link&type=grpc&serviceName=@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray#%F0%9F%87%B7%F0%9F%87%BA%20%D8%B1%D9%88%D8%B3%DB%8C%D9%87%20:%20@KralVPN
		vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray`)

		fmt.Println(len(links), links)

		if len(links) != 9 {
			t.Error("tf is going on")
		}

		if !slices.Contains(links, "vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray") &&
			!slices.Contains(links, "vless://1cb85996-a546-4b7a-bc5a-fefc35da7939@mkh.arv2ray.host:2083?security=tls&sni=gr.ArV2ray.link&type=grpc&serviceName=@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray#%F0%9F%87%B7%F0%9F%87%BA%20%D8%B1%D9%88%D8%B3%DB%8C%D9%87%20:%20@KralVPN") {
			t.Error("the link finder works shitty!!")
		}
	})
}

func TestFindVmessLinksInText(t *testing.T) {
	vmess := `💯 کانفیگ Vmess
	📍 لوکیشن : 🇮🇷 ایران به خارج (فوروارد ترافیک)
	
	🔗 Link : vmess://eydhZGQnOiAnZWR1Lm1yc3BvdC50aycsICdhaWQnOiAnMCcsICdhbHBuJzogJ2gyLGh0dHAvMS4xJywgJ2ZwJzogJ3JhbmRvbWl6ZWQnLCAnaG9zdCc6ICdlZHUubXJzcG90LnRrJywgJ2lkJzogJzZhYzdkM2VlLTgyMzEtNDM2Ny1hZTFhLTI4ZWNlZDQzY2ZiMScsICduZXQnOiAndGNwJywgJ3BhdGgnOiAnL3RjcDE5MzI2JywgJ3BvcnQnOiAnMTkzMjYnLCAncHMnOiAn8J+HrvCfh7cg2KfbjNix2KfZhiDYqNmHINiu2KfYsdisICjZgdmI2LHZiNin2LHYryDYqtix2KfZgduM2qkpIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJ2VkdS5tcnNwb3QudGsnLCAndGxzJzogJ3RscycsICd0eXBlJzogJ2h0dHAnLCAndic6ICcyJ30=
	
	⚙ isp : Web Dadeh Paydar Co (Ltd)
	📍 @KralVPN`

	links := skraper.FindVmessLinkInText(vmess)

	if !slices.Contains(links, "vmess://eydhZGQnOiAnZWR1Lm1yc3BvdC50aycsICdhaWQnOiAnMCcsICdhbHBuJzogJ2gyLGh0dHAvMS4xJywgJ2ZwJzogJ3JhbmRvbWl6ZWQnLCAnaG9zdCc6ICdlZHUubXJzcG90LnRrJywgJ2lkJzogJzZhYzdkM2VlLTgyMzEtNDM2Ny1hZTFhLTI4ZWNlZDQzY2ZiMScsICduZXQnOiAndGNwJywgJ3BhdGgnOiAnL3RjcDE5MzI2JywgJ3BvcnQnOiAnMTkzMjYnLCAncHMnOiAn8J+HrvCfh7cg2KfbjNix2KfZhiDYqNmHINiu2KfYsdisICjZgdmI2LHZiNin2LHYryDYqtix2KfZgduM2qkpIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJ2VkdS5tcnNwb3QudGsnLCAndGxzJzogJ3RscycsICd0eXBlJzogJ2h0dHAnLCAndic6ICcyJ30=") {
		fmt.Println(links)
		t.Error("the link finder works shitty")
	}
}
