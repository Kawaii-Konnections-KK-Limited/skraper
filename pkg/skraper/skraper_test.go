package skraper_test

import (
	"fmt"
	"skraper/pkg/skraper"
	"testing"
)

func TestFindLinksInText(t *testing.T) {
	t.Run("findtext", func(t *testing.T) {
		fmt.Println(skraper.FindLinksInText("کانفیگ اهدایی:
		@iran_ray
		
		
		لوکیشن هلند برای دوستانی که درخواست کرده بودند
		
		
		محدودیت حجم: 1 ترابایت
		محدودیت روز:  ندارد
		
		vless://iran-ray-iran-ray@iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.iran-ray.top:443?security=none&encryption=none&host=fast.com&headerType=http&type=tcp#-Telegram%3A%40iran-ray"))
	})

}