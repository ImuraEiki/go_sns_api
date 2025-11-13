package infrastructure

import (
	"log"
	"my-gin-app/internal/domain"

	"gorm.io/gorm"
)

func SeedInitialData(db *gorm.DB) {
	var postCounts int64
	db.Model(&domain.Post{}).Count(&postCounts)
	if postCounts == 0 {
		// 初期値がまだ入っていない場合のみ挿入
		initialPosts := []domain.Post{
			{Content: "APIから返却してます", Likes: 8, UserId: 9},
			{Content: "転倒狭い持ってるフェミニスト。", Likes: 1, UserId: 1},
			{Content: "奨励します呼ぶパーセント大統領リンク主人。", Likes: 1, UserId: 1},
			{Content: "じぶんのシュガー極端な今日。", Likes: 1, UserId: 1},
			{Content: "呼ぶパイオニア犯罪者雪叔父。", Likes: 1, UserId: 1},
			{Content: "ハンマー奨励します私サワー転倒緩む。", Likes: 1, UserId: 1},
			{Content: "タワー敵ダニ部隊。", Likes: 1, UserId: 1},
			{Content: "主人ブレーキ午前インチ血まみれの。", Likes: 2, UserId: 2},
			{Content: "カラムクルー改善厳しい残る欠乏。", Likes: 2, UserId: 2},
			{Content: "持っていましたコーナー催眠術デフォルトメニューバナー販売。", Likes: 2, UserId: 2},
			{Content: "普通のジャム今日隠す衝突障害運。", Likes: 2, UserId: 4},
			{Content: "トレーナー呼ぶ状況見出しヒット脊椎倫理リニア。", Likes: 4, UserId: 4},
			{Content: "ヒット軸行進ニューススペル。", Likes: 4, UserId: 39},
			{Content: "楽しんでハードウェア錯覚ヒール省略人形。", Likes: 4, UserId: 4},
			{Content: "意図メニュー欠乏偏差サンプル。", Likes: 4, UserId: 92},
			{Content: "残る欠乏偏差。", Likes: 4, UserId: 5},
			{Content: "鉱山プラスチックあなた自身日曜日パイオニア。", Likes: 5, UserId: 5},
			{Content: "午前敵対的な尊敬する教授バケツ血まみれの立派なスマッシュ。", Likes: 5, UserId: 5},
			{Content: "ブラケットサラダ電話タワー品質リフト指名。", Likes: 5, UserId: 9},
			{Content: "欠乏符号パーセント腐った柔らかい。", Likes: 6, UserId: 6},
			{Content: "見落とすログフェミニスト尿。", Likes: 44, UserId: 6},
			{Content: "午前サワー普通の発生するハードウェア憲法。", Likes: 6, UserId: 63},
			{Content: "軸極端な人形本質的なリハビリフレーム。", Likes: 6, UserId: 6},
			{Content: "キャビン花嫁索引ブレーキ動物。", Likes: 6, UserId: 3},
			{Content: "立派な厳しい分割スペル証言する溝。", Likes: 3, UserId: 64},
			{Content: "ダイヤモンド彼女中央普通の戦略的。", Likes: 8, UserId: 8},
			{Content: "じぶんのヘアハードウェアピック溝デフォルト省略。", Likes: 65, UserId: 91},
			{Content: "コーナーデッドバナー。", Likes: 55, UserId: 15},
			{Content: "文言スキームない品質。", Likes: 15, UserId: 8},
			{Content: "学生高い主婦評議会職人雪。", Likes: 8, UserId: 8},
		}
		if err := db.Create(&initialPosts).Error; err != nil {
			log.Printf("❌ 初期Postデータ投入に失敗: %v", err)
		} else {
			log.Println("✅ 初期Postデータ投入完了")
		}
	}
	var userCounts int64
	db.Model(&domain.User{}).Count(&userCounts)
	if userCounts == 0 {
		// 初期値がまだ入っていない場合のみ挿入
		initialUsers := []domain.User{
			{Name: "山田 晃", Email: "satokana@sato.com"}, {Name: "岡田 充", Email: "yutasato@sato.jp"}, {Name: "山口 聡太郎", Email: "umori@sato.com"}, {Name: "山本 くみ子", Email: "gotoyoko@saito.jp"}, {Name: "岡田 里佳", Email: "mai63@yahoo.com"}, {Name: "石川 七夏", Email: "atsushisuzuki@gmail.com"}, {Name: "中村 翔太", Email: "asukatakahashi@hotmail.com"}, {Name: "吉田 幹", Email: "mituruhashimoto@hayashi.jp"}, {Name: "藤原 直人", Email: "ukobayashi@yahoo.com"}, {Name: "村上 幹", Email: "ishiimituru@watanabe.jp"}, {Name: "渡辺 健一", Email: "kyosuke86@yahoo.com"}, {Name: "前田 治", Email: "akemitakahashi@yamamoto.com"}, {Name: "加藤 あすか", Email: "yumiko03@sasaki.com"}, {Name: "渡辺 翼", Email: "sotaro18@gmail.com"}, {Name: "石川 知実", Email: "yuiyamamoto@aoki.net"}, {Name: "清水 花子", Email: "naoki90@fujita.jp"}, {Name: "山本 香織", Email: "kumikookada@nakagawa.com"}, {Name: "藤井 真綾", Email: "momoko84@matsumoto.net"}, {Name: "中村 修平", Email: "cwatanabe@yahoo.com"}, {Name: "山本 里佳", Email: "satokenichi@gmail.com"}, {Name: "中島 充", Email: "shimizutaro@yahoo.com"}, {Name: "吉田 浩", Email: "satoshota@yahoo.com"}, {Name: "藤原 結衣", Email: "ainoue@kobayashi.com"}, {Name: "井上 洋介", Email: "akirayoshida@suzuki.com"}, {Name: "阿部 零", Email: "osamu38@yahoo.com"}, {Name: "佐藤 舞", Email: "yamamotohideki@yahoo.com"}, {Name: "岡本 美加子", Email: "hsaito@tanaka.jp"}, {Name: "井上 香織", Email: "ksaito@ito.org"}, {Name: "渡辺 加奈", Email: "lsuzuki@suzuki.org"}, {Name: "高橋 くみ子", Email: "akira59@mori.org"}, {Name: "山下 智也", Email: "takahashiyosuke@gmail.com"}, {Name: "斉藤 明美", Email: "murakamimanabu@yahoo.com"}, {Name: "山田 智也", Email: "fujiwaraakemi@okamoto.org"}, {Name: "石川 和也", Email: "rikatakahashi@fukuda.jp"}, {Name: "中島 裕美子", Email: "asuka36@yamamoto.com"}, {Name: "松本 さゆり", Email: "xnakamura@hotmail.com"}, {Name: "山本 涼平", Email: "yokoito@hotmail.com"}, {Name: "阿部 舞", Email: "matsumotoyasuhiro@okamoto.net"}, {Name: "佐藤 充", Email: "rikayamada@hotmail.com"}, {Name: "後藤 淳", Email: "yamashitayuta@hotmail.com"}, {Name: "村上 直人", Email: "yamamotorika@gmail.com"}, {Name: "藤田 直子", Email: "hidekigoto@gmail.com"}, {Name: "山下 結衣", Email: "suzukiasuka@tanaka.net"}, {Name: "佐藤 くみ子", Email: "naoko40@kimura.jp"}, {Name: "森 直樹", Email: "kumikokimura@suzuki.jp"}, {Name: "藤田 里佳", Email: "tishikawa@yahoo.com"}, {Name: "高橋 涼平", Email: "yukisato@nakajima.jp"}, {Name: "林 直子", Email: "eishii@gmail.com"}, {Name: "山口 篤司", Email: "satomiito@tanaka.jp"}, {Name: "清水 聡太郎", Email: "kyoshida@hotmail.com"},
		}
		if err := db.Create(&initialUsers).Error; err != nil {
			log.Printf("❌ 初期Userデータ投入に失敗: %v", err)
		} else {
			log.Println("✅ 初期Userデータ投入完了")
		}
	}
	// var commentCounts int64
	// db.Model(&domain.Comment{}).Count(&commentCounts)
	// if commentCounts == 0 {
	// initialComments := []domain.Comment{}
	// if err := db.Create(&initialComments).Error; err != nil {
	// 	log.Printf("❌ 初期Commentデータ投入に失敗: %v", err)
	// } else {
	// 	log.Println("✅ 初期Commentデータ投入完了")
	// }
	// }
	// var followingCounts int64
	// db.Model(&domain.Following{}).Count(&followingCounts)
	// if followingCounts == 0 {
	// initialFollowings := []domain.Following{}
	// if err := db.Create(&initialFollowings).Error; err != nil {
	// 	log.Printf("❌ 初期Followingデータ投入に失敗: %v", err)
	// } else {
	// 	log.Println("✅ 初期Followingデータ投入完了")
	// }
	// }
}
