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
		vmess://eydhZGQnOiAnZWR1Lm1yc3BvdC50aycsICdhaWQnOiAnMCcsICdhbHBuJzogJ2gyLGh0dHAvMS4xJywgJ2ZwJzogJ3JhbmRvbWl6ZWQnLCAnaG9zdCc6ICdlZHUubXJzcG90LnRrJywgJ2lkJzogJzZhYzdkM2VlLTgyMzEtNDM2Ny1hZTFhLTI4ZWNlZDQzY2ZiMScsICduZXQnOiAndGNwJywgJ3BhdGgnOiAnL3RjcDE5MzI2JywgJ3BvcnQnOiAnMTkzMjYnLCAncHMnOiAn8J+HrvCfh7cg2KfbjNix2KfZhiDYqNmHINiu2KfYsdisICjZgdmI2LHZiNin2LHYryDYqtix2KfZgduM2qkpIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJ2VkdS5tcnNwb3QudGsnLCAndGxzJzogJ3RscycsICd0eXBlJzogJ2h0dHAnLCAndic6ICcyJ30=
		vless://1cb85996-a546-4b7a-bc5a-fefc35da7939@mkh.arv2ray.host:2083?security=tls&sni=gr.ArV2ray.link&type=grpc&serviceName=@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray,@ArV2ray#%F0%9F%87%B7%F0%9F%87%BA%20%D8%B1%D9%88%D8%B3%DB%8C%D9%87%20:%20@KralVPN
		vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray`)

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
