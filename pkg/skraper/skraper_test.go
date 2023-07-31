package skraper_test

import (
	"skraper/pkg/skraper"
	"testing"

	"golang.org/x/exp/slices"
)

func TestFindLinksInText(t *testing.T) {
	t.Run("findtext", func(t *testing.T) {
		links := skraper.FindLinksInText(`کانفیگ اهدایی:
		@iran_ray
		
		
		لوکیشن هلند برای دوستانی که درخواست کرده بودند
		
		
		محدودیت حجم: 1 ترابایت
		محدودیت روز:  ندارد
		
		vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray`)

		if !slices.Contains(links, "vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray") {
			t.Error("the link finder works shitty!!")
		}
	})
}
