package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	for _, unit := range []struct {
		orig string
		want string
	}{
		{"", ""},
		{"n9110", "n9110"},
		{"18ntrd052", "18ntrd052"},
		{"118abp077", "118abp077"},
		{"5abp077", "5abp077"},
		{"T-28621", "T-28621"},
		{"h_346rebd655tk2", "h_346rebd655tk2"},
		{"200gana-1350", "200gana-1350"},
		{"ssis00123", "ssis00123"},
		{"ABP-030", "ABP-030"},
		{"MARRA-A030", "MARRA-A030"},
		{"MARRA-A030-C", "MARRA-A030"},
		{"ABP-030-C", "ABP-030"},
		{"ABP-030C", "ABP-030"},
		{"ABP-030A", "ABP-030"},
		{"ABP-030B", "ABP-030"},
		{"ABP-030-A", "ABP-030"},
		{"ABP-030-C.mp4", "ABP-030"},
		{"ABP-358_C.mkv", "ABP-358"},
		{"[]ABP-358_C.mkv", "ABP-358"},
		{"[22sht.me]ABP-358_C.mkv", "ABP-358"},
		{"[98t.tv]vema-181-4k-C.mp4", "vema-181"},
		{"[98t.tv]vema-180-4k-C", "vema-180"},
		{"[98t.tv]vema-180-4k-C.mp4", "vema-180"},
		{"ABP-030-C-c_c-C-Cd1-cd4.mp4", "ABP-030"},
		{"rctd-460ch.mp4", "rctd-460"},
		{"rctd-460-ch.mp4", "rctd-460"},
		{"rctd-460ch-ch.mp4", "rctd-460"},
		{"rctd-461-C-cD4.mp4", "rctd-461"},
		{"rctd-461-cd3.mp4", "rctd-461"},
		{"rctd-461-Cd3-C.mp4", "rctd-461"},
		{"rctd-461_Cd39-C.mp4", "rctd-461"},
		{"FC2-PPV-123456", "FC2-123456"},
		{"FC2PPV-123456", "FC2-123456"},
		{"FC2PPV-123456-1", "FC2-123456"},
		{"FC2-123456-1", "FC2-123456"},
		{"FC2PPV_123456", "FC2-123456"},
		{"FC2_PPV_123456", "FC2-123456"},
		{"FC2-PPV_123456", "FC2-123456"},
		{"FC2-PPV-123456-C.mp4", "FC2-123456"},
		{"cllshuai@www.sexinsex.net@[中文字幕]FAX-146剧场版-和大嫂偷情、迷奸、强奸等4个故事", "FAX-146"},
		{"SHKD-474 KATAGIRI ERIRIKA 카타기리 에리리카 - FHD 1080P", "SHKD-474"},
		{"【WRE】SHKD-474 shezhangmi", "SHKD-474"},
		{"[Attackers] SHKD-474  社長秘書美畜同好今 輪姦標的 (Eririka Kitagari).avi", "SHKD-474"},
		{"HD_GS-333", "GS-333"},
		{"FHD-MXGS-247-C", "MXGS-247"},
		{"17.ebod-185", "ebod-185"},
		{"17.ebod-185-C", "ebod-185"},
		{"【TXH】EBOD-185 E-BODY -te", "EBOD-185"},
		{"nike@EBOD-185 DVD", "EBOD-185"},
		{"HDD-697-C-dvd.mp4", "HDD-697"},
		{"HND-697-C-dvd.mp4", "HND-697"},
		{"DVDES-697-C-dvd.mp4", "DVDES-697"},
		{"MXGS-697.HD.mp4", "MXGS-697"},
		{"MXGS-697-AVI", "MXGS-697"},
		{"MXGS-697_CAVI", "MXGS-697"},
		{"MXGS-697-MP4", "MXGS-697"},
		{"MXGS-797,.avi", "MXGS-797"},
		{"avop-208-hhbhd.mp4", "avop-208"},
		{"(無修正-流出) MXGS-247.mp4", "MXGS-247"},
		{"[Uncensored] SDDE-618", "SDDE-618"},
		{"SDDE-618 NEWS .mp4", "SDDE-618"},
		{"ssis00123.mp4", "ssis00123"},
		{"SDDE-625_uncensored_C", "SDDE-625"},
		{"SDDE-625_uncensored_C.mp4", "SDDE-625"},
		{"SDDE-625_uncensored_leak_C.mp4", "SDDE-625"},
		{"SDDE-625_uncensored_leak_C_cd1.mp4", "SDDE-625"},
		{"GIGL-677_4K.mp4", "GIGL-677"},
		{"GIGL-677_2K_h265.mp4", "GIGL-677"},
		{"GIGL-677_4K60FPS.mp4", "GIGL-677"},
		{"093021_539-FHD.mkv", "093021_539"},
		{"093021_539-480p.mkv", "093021_539"},
		{"093021_539-1080pFHD.mkv", "093021_539"},
		{"093021_539-2160pHD.mkv", "093021_539"},
		{"SSIS-329_60FPS", "SSIS-329"},
		{"SSIS-329-C_60FPS", "SSIS-329"},
		{"SSIS-329-C_1080P30FPS", "SSIS-329"},
		{"SSIS-329-C_1080P30FPSFHDx264", "SSIS-329"},
		{"hhd800.com@HUNTB-269", "HUNTB-269"},
		{"sbw99.cc@iesp-653-4K.mp4", "iesp-653"},
		{"jav20s8.com@GIGL-677.mp4", "GIGL-677"},
		{"jav20s8.com@GIGL-677_4K.mp4", "GIGL-677"},
		{"133ARA-030你好.mp4", "133ARA-030"},
		{"133ARA-030 你好.mp4", "133ARA-030"},
		{"133ARA-030 hello there", "133ARA-030"},
		{"133ARA-030 hello there.mp4", "133ARA-030"},
		{"133ARA-030 - hello there.mp4", "133ARA-030"},
		{"133ARA-030-C 你好.mp4", "133ARA-030"},
		{"n9930 你好.mp4", "n9930"},
		{"133ARA-030-C - 你好.mp4", "133ARA-030"},
		{"test.xxx@133ARA-030 你好", "133ARA-030"},
		{"test.xxx@133ARA-030 你好.mp4", "133ARA-030"},
		{"Tokyo Hot n9001 FHD.mp4", "n9001"},
		{"TokyoHot-n1287-HD .mp4", "n1287"},
		{"caribean-020317_001.mp4", "020317_001"},
		{"caribbean-020317_001.mp4", "020317_001"},
		{"caribeancom-020317_001.mp4", "020317_001"},
		{"caribbeancom-020317_001.mp4", "020317_001"},
		{"020317_001-caribbeancom.mp4", "020317_001"},
		{"020317_001-caribbean 你好.mp4", "020317_001"},
		{"020317_001-caribbean-fhd 你好.mp4", "020317_001"},
		{"020317_001-caribbean.mp4", "020317_001"},
		{"020317_001-carib-whole_hd1.mp4", "020317_001"},
		{"020317_001-carib-whole_fhd1.mp4", "020317_001"},
		{"020317_001-carib_sd1.mp4", "020317_001"},
		{"carib-020317_001.mp4", "020317_001"},
		{"020317-001-1pondo.mp4", "020317-001"},
		{"020317-001-pondo.mp4", "020317-001"},
		{"020317-001-pond.mp4", "020317-001"},
		{"1pondo-020317-001.mp4", "020317-001"},
		{"_1PONDO_020317-001.mp4", "020317-001"},
		{"pond-020317-001.mp4", "020317-001"},
		{"pondo-020317-001.mp4", "020317-001"},
		{"4102-023-CD2.iso", "4102-023"},
		{"heydouga-4102-023.mp4", "heydouga-4102-023"},
		{"heydouga-4102-023-CD2.iso", "heydouga-4102-023"},
		{"xxx-av-1789.mp4", "xxx-av-1789"},
		{"xxx-av-1789-cd1.mp4", "xxx-av-1789"},
		{"xxx-av-1789-C.mp4", "xxx-av-1789"},
		{"10musume-020317_01-CD2.iso", "020317_01"},
		{"HND-621 Jinguuji Nao Cry Unintentionally- HD", "HND-621"},
		{"[ThZu.Cc]hnd-621", "hnd-621"},
		{"[CAND-1196]恋するスチューピッド", "CAND-1196"},
		{"[98t.tv][98t.tv]300MAAN-791", "300MAAN-791"},
		{"(HND-620) 絶対にナマで連射させてくれる連続中出しソープ あいだ飛鳥", "HND-620"},
		{"第一會所新片@SIS001@(本中)(HND-621)気持ち良すぎて思わず叫んじゃうごめんなさいGスポットずーっと腰振り回し続けるイクイク騎乗位中出し_神宮寺ナオ", "HND-621"},
		{"第一會所新片@SIS001@(本中)(HND-620)絶対にナマで連射させてくれる連続中出しソープ_あいだ飛鳥", "HND-620"},
		{"HND-620絶対にナマで連射させてくれる連続中出しソープ_あいだ飛鳥", "HND-620"},
		{"絶対にナマで連射させてくれる連続中出しソープ_あいだ飛鳥 HND-620", "HND-620"},
		{"kb1234 絶対にナマで連射させてくれる連続中出しソープ_あいだ飛鳥", "kb1234"},
		{"絶対にナマで連射させてくれる連続中出しソープ_あいだ飛鳥 kb1234", "kb1234"},
		{"hnd-993ch字幕", "hnd-993"},
		{"japlib.top_HND-993.mp4", "HND-993"},
		{"bbs2048@HND-972-SD", "HND-972"},
		{"zzpp06.com@n1666", "n1666"},
		{"[psk.la]Carib-080520-001-FHD", "080520-001"},
		{"Carib 080520-001.mp4", "080520-001"},
		{"[ThZu.Cc]080520-001-carib-720p", "080520-001"},
		{"MCBD-18 Merci Beaucoup 18 Merci Beaucoup Beach Fuck Special 3Hrs 14Girls (Blu-ray) {18iso.com} [1080p].mp4", "MCBD-18"},
		{"javcn.net_MCBD-18-H265.mp4", "MCBD-18"},
		{"MIDV-111_X1080X.mp4", "MIDV-111"},
		{"MIDV-111-C_X1080X.mp4", "MIDV-111"},
		{"hhd800.com@MIDV-111-C_X1080X.mp4", "MIDV-111"},
	} {
		assert.Equal(t, unit.want, Trim(unit.orig), unit.orig)
	}
}

func TestIsUncensored(t *testing.T) {
	for _, unit := range []struct {
		orig string
		want bool
	}{
		{"ABP-030", false},
		{"ssis00123", false},
		{"133ARA-030", false},
		{"123456_789", true},
		{"123456-789", true},
		{"123456-01", true},
		{"xxx-av-1789", true},
		{"xxx-av_1789", true},
		{"heydouga-1789-233", true},
		{"heydouga_1789-233", true},
		{"heyzo-1342", true},
		{"heyzo_1342", true},
		{"n1342", true},
		{"kb1342", true},
	} {
		assert.Equal(t, unit.want, IsUncensored(unit.orig), unit.orig)
	}
}

func TestIsFC2(t *testing.T) {
	for _, unit := range []struct {
		orig string
		want bool
	}{
		{"FC2-738573", true},
		{"FC2_738573", true},
		{"FC2-PPV-738573", true},
		{"FC2PPV-738573", true},
		{"FC2PPV738573", true},
		{"ABC-123", false},
		{"fc2996-123", false},
	} {
		assert.Equal(t, unit.want, IsFC2(unit.orig), unit.orig)
	}
}

func TestIsSpecial(t *testing.T) {
	for _, unit := range []struct {
		orig string
		want bool
	}{
		{"ABP-133", false},
		{"FC2-738573", true},
		{"FC2_738573", true},
		{"133ARA-030", false},
		{"123456_789", true},
		{"123456-789", true},
		{"FC2-PPV-738573", true},
		{"xxx-av-1789", true},
		{"xxx-av_1789", true},
		{"heydouga-1789-233", true},
		{"heydouga_1789-233", true},
		{"heyzo-1342", true},
		{"heyzo_1342", true},
		{"pcolle-1332342", true},
	} {
		assert.Equal(t, unit.want, IsSpecial(unit.orig), unit.orig)
	}
}

func TestRequireFaceDetection(t *testing.T) {
	for _, unit := range []struct {
		orig string
		want bool
	}{
		{"ABP-030", false},
		{"ssis00123", false},
		{"SIRO-030", true},
		{"133ARA-030", true},
		{"FC2-738573", true},
		{"123456_789", true},
		{"123456-01", true},
		{"xxx-av-1789", true},
		{"heydouga-1789-233", true},
		{"heyzo-1342", true},
		{"n1342", true},
		{"kb1342", true},
		{"gcolle-847256", true},
		{"pcolle-14491760b933a35cfab", true},
		{"gyutto-254274", true},
		{"getchu-4041236", true},
		{"orec-062", true},
		{"orec062", true},
		{"orecw-062", false},
		{"LAFBD-41", false},
		{"PRED-314", false},
	} {
		assert.Equal(t, unit.want, RequireFaceDetection(unit.orig), unit.orig)
	}
}
