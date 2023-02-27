package jpostcode

import (
	"github.com/syumai/go-jpostcode/internal/address"
	"github.com/syumai/go-jpostcode/internal/data"
)

type varAdapter struct{}

func (a *varAdapter) SearchAddressesFromPostCode(postCode string) ([]*Address, error) {
	first := postCode[:3]
	second := postCode[3:]
	switch first {
	case "001":
		return fromDataAddresses(data.Addresses001, second)
	case "002":
		return fromDataAddresses(data.Addresses002, second)
	case "003":
		return fromDataAddresses(data.Addresses003, second)
	case "004":
		return fromDataAddresses(data.Addresses004, second)
	case "005":
		return fromDataAddresses(data.Addresses005, second)
	case "006":
		return fromDataAddresses(data.Addresses006, second)
	case "007":
		return fromDataAddresses(data.Addresses007, second)
	case "010":
		return fromDataAddresses(data.Addresses010, second)
	case "011":
		return fromDataAddresses(data.Addresses011, second)
	case "012":
		return fromDataAddresses(data.Addresses012, second)
	case "013":
		return fromDataAddresses(data.Addresses013, second)
	case "014":
		return fromDataAddresses(data.Addresses014, second)
	case "015":
		return fromDataAddresses(data.Addresses015, second)
	case "016":
		return fromDataAddresses(data.Addresses016, second)
	case "017":
		return fromDataAddresses(data.Addresses017, second)
	case "018":
		return fromDataAddresses(data.Addresses018, second)
	case "019":
		return fromDataAddresses(data.Addresses019, second)
	case "020":
		return fromDataAddresses(data.Addresses020, second)
	case "021":
		return fromDataAddresses(data.Addresses021, second)
	case "022":
		return fromDataAddresses(data.Addresses022, second)
	case "023":
		return fromDataAddresses(data.Addresses023, second)
	case "024":
		return fromDataAddresses(data.Addresses024, second)
	case "025":
		return fromDataAddresses(data.Addresses025, second)
	case "026":
		return fromDataAddresses(data.Addresses026, second)
	case "027":
		return fromDataAddresses(data.Addresses027, second)
	case "028":
		return fromDataAddresses(data.Addresses028, second)
	case "029":
		return fromDataAddresses(data.Addresses029, second)
	case "030":
		return fromDataAddresses(data.Addresses030, second)
	case "031":
		return fromDataAddresses(data.Addresses031, second)
	case "033":
		return fromDataAddresses(data.Addresses033, second)
	case "034":
		return fromDataAddresses(data.Addresses034, second)
	case "035":
		return fromDataAddresses(data.Addresses035, second)
	case "036":
		return fromDataAddresses(data.Addresses036, second)
	case "037":
		return fromDataAddresses(data.Addresses037, second)
	case "038":
		return fromDataAddresses(data.Addresses038, second)
	case "039":
		return fromDataAddresses(data.Addresses039, second)
	case "040":
		return fromDataAddresses(data.Addresses040, second)
	case "041":
		return fromDataAddresses(data.Addresses041, second)
	case "042":
		return fromDataAddresses(data.Addresses042, second)
	case "043":
		return fromDataAddresses(data.Addresses043, second)
	case "044":
		return fromDataAddresses(data.Addresses044, second)
	case "045":
		return fromDataAddresses(data.Addresses045, second)
	case "046":
		return fromDataAddresses(data.Addresses046, second)
	case "047":
		return fromDataAddresses(data.Addresses047, second)
	case "048":
		return fromDataAddresses(data.Addresses048, second)
	case "049":
		return fromDataAddresses(data.Addresses049, second)
	case "050":
		return fromDataAddresses(data.Addresses050, second)
	case "051":
		return fromDataAddresses(data.Addresses051, second)
	case "052":
		return fromDataAddresses(data.Addresses052, second)
	case "053":
		return fromDataAddresses(data.Addresses053, second)
	case "054":
		return fromDataAddresses(data.Addresses054, second)
	case "055":
		return fromDataAddresses(data.Addresses055, second)
	case "056":
		return fromDataAddresses(data.Addresses056, second)
	case "057":
		return fromDataAddresses(data.Addresses057, second)
	case "058":
		return fromDataAddresses(data.Addresses058, second)
	case "059":
		return fromDataAddresses(data.Addresses059, second)
	case "060":
		return fromDataAddresses(data.Addresses060, second)
	case "061":
		return fromDataAddresses(data.Addresses061, second)
	case "062":
		return fromDataAddresses(data.Addresses062, second)
	case "063":
		return fromDataAddresses(data.Addresses063, second)
	case "064":
		return fromDataAddresses(data.Addresses064, second)
	case "065":
		return fromDataAddresses(data.Addresses065, second)
	case "066":
		return fromDataAddresses(data.Addresses066, second)
	case "067":
		return fromDataAddresses(data.Addresses067, second)
	case "068":
		return fromDataAddresses(data.Addresses068, second)
	case "069":
		return fromDataAddresses(data.Addresses069, second)
	case "070":
		return fromDataAddresses(data.Addresses070, second)
	case "071":
		return fromDataAddresses(data.Addresses071, second)
	case "072":
		return fromDataAddresses(data.Addresses072, second)
	case "073":
		return fromDataAddresses(data.Addresses073, second)
	case "074":
		return fromDataAddresses(data.Addresses074, second)
	case "075":
		return fromDataAddresses(data.Addresses075, second)
	case "076":
		return fromDataAddresses(data.Addresses076, second)
	case "077":
		return fromDataAddresses(data.Addresses077, second)
	case "078":
		return fromDataAddresses(data.Addresses078, second)
	case "079":
		return fromDataAddresses(data.Addresses079, second)
	case "080":
		return fromDataAddresses(data.Addresses080, second)
	case "081":
		return fromDataAddresses(data.Addresses081, second)
	case "082":
		return fromDataAddresses(data.Addresses082, second)
	case "083":
		return fromDataAddresses(data.Addresses083, second)
	case "084":
		return fromDataAddresses(data.Addresses084, second)
	case "085":
		return fromDataAddresses(data.Addresses085, second)
	case "086":
		return fromDataAddresses(data.Addresses086, second)
	case "087":
		return fromDataAddresses(data.Addresses087, second)
	case "088":
		return fromDataAddresses(data.Addresses088, second)
	case "089":
		return fromDataAddresses(data.Addresses089, second)
	case "090":
		return fromDataAddresses(data.Addresses090, second)
	case "091":
		return fromDataAddresses(data.Addresses091, second)
	case "092":
		return fromDataAddresses(data.Addresses092, second)
	case "093":
		return fromDataAddresses(data.Addresses093, second)
	case "094":
		return fromDataAddresses(data.Addresses094, second)
	case "095":
		return fromDataAddresses(data.Addresses095, second)
	case "096":
		return fromDataAddresses(data.Addresses096, second)
	case "097":
		return fromDataAddresses(data.Addresses097, second)
	case "098":
		return fromDataAddresses(data.Addresses098, second)
	case "099":
		return fromDataAddresses(data.Addresses099, second)
	case "100":
		return fromDataAddresses(data.Addresses100, second)
	case "101":
		return fromDataAddresses(data.Addresses101, second)
	case "102":
		return fromDataAddresses(data.Addresses102, second)
	case "103":
		return fromDataAddresses(data.Addresses103, second)
	case "104":
		return fromDataAddresses(data.Addresses104, second)
	case "105":
		return fromDataAddresses(data.Addresses105, second)
	case "106":
		return fromDataAddresses(data.Addresses106, second)
	case "107":
		return fromDataAddresses(data.Addresses107, second)
	case "108":
		return fromDataAddresses(data.Addresses108, second)
	case "109":
		return fromDataAddresses(data.Addresses109, second)
	case "110":
		return fromDataAddresses(data.Addresses110, second)
	case "111":
		return fromDataAddresses(data.Addresses111, second)
	case "112":
		return fromDataAddresses(data.Addresses112, second)
	case "113":
		return fromDataAddresses(data.Addresses113, second)
	case "114":
		return fromDataAddresses(data.Addresses114, second)
	case "115":
		return fromDataAddresses(data.Addresses115, second)
	case "116":
		return fromDataAddresses(data.Addresses116, second)
	case "120":
		return fromDataAddresses(data.Addresses120, second)
	case "121":
		return fromDataAddresses(data.Addresses121, second)
	case "123":
		return fromDataAddresses(data.Addresses123, second)
	case "124":
		return fromDataAddresses(data.Addresses124, second)
	case "125":
		return fromDataAddresses(data.Addresses125, second)
	case "130":
		return fromDataAddresses(data.Addresses130, second)
	case "131":
		return fromDataAddresses(data.Addresses131, second)
	case "132":
		return fromDataAddresses(data.Addresses132, second)
	case "133":
		return fromDataAddresses(data.Addresses133, second)
	case "134":
		return fromDataAddresses(data.Addresses134, second)
	case "135":
		return fromDataAddresses(data.Addresses135, second)
	case "136":
		return fromDataAddresses(data.Addresses136, second)
	case "137":
		return fromDataAddresses(data.Addresses137, second)
	case "140":
		return fromDataAddresses(data.Addresses140, second)
	case "141":
		return fromDataAddresses(data.Addresses141, second)
	case "142":
		return fromDataAddresses(data.Addresses142, second)
	case "143":
		return fromDataAddresses(data.Addresses143, second)
	case "144":
		return fromDataAddresses(data.Addresses144, second)
	case "145":
		return fromDataAddresses(data.Addresses145, second)
	case "146":
		return fromDataAddresses(data.Addresses146, second)
	case "150":
		return fromDataAddresses(data.Addresses150, second)
	case "151":
		return fromDataAddresses(data.Addresses151, second)
	case "152":
		return fromDataAddresses(data.Addresses152, second)
	case "153":
		return fromDataAddresses(data.Addresses153, second)
	case "154":
		return fromDataAddresses(data.Addresses154, second)
	case "155":
		return fromDataAddresses(data.Addresses155, second)
	case "156":
		return fromDataAddresses(data.Addresses156, second)
	case "157":
		return fromDataAddresses(data.Addresses157, second)
	case "158":
		return fromDataAddresses(data.Addresses158, second)
	case "160":
		return fromDataAddresses(data.Addresses160, second)
	case "161":
		return fromDataAddresses(data.Addresses161, second)
	case "162":
		return fromDataAddresses(data.Addresses162, second)
	case "163":
		return fromDataAddresses(data.Addresses163, second)
	case "164":
		return fromDataAddresses(data.Addresses164, second)
	case "165":
		return fromDataAddresses(data.Addresses165, second)
	case "166":
		return fromDataAddresses(data.Addresses166, second)
	case "167":
		return fromDataAddresses(data.Addresses167, second)
	case "168":
		return fromDataAddresses(data.Addresses168, second)
	case "169":
		return fromDataAddresses(data.Addresses169, second)
	case "170":
		return fromDataAddresses(data.Addresses170, second)
	case "171":
		return fromDataAddresses(data.Addresses171, second)
	case "173":
		return fromDataAddresses(data.Addresses173, second)
	case "174":
		return fromDataAddresses(data.Addresses174, second)
	case "175":
		return fromDataAddresses(data.Addresses175, second)
	case "176":
		return fromDataAddresses(data.Addresses176, second)
	case "177":
		return fromDataAddresses(data.Addresses177, second)
	case "178":
		return fromDataAddresses(data.Addresses178, second)
	case "179":
		return fromDataAddresses(data.Addresses179, second)
	case "180":
		return fromDataAddresses(data.Addresses180, second)
	case "181":
		return fromDataAddresses(data.Addresses181, second)
	case "182":
		return fromDataAddresses(data.Addresses182, second)
	case "183":
		return fromDataAddresses(data.Addresses183, second)
	case "184":
		return fromDataAddresses(data.Addresses184, second)
	case "185":
		return fromDataAddresses(data.Addresses185, second)
	case "186":
		return fromDataAddresses(data.Addresses186, second)
	case "187":
		return fromDataAddresses(data.Addresses187, second)
	case "188":
		return fromDataAddresses(data.Addresses188, second)
	case "189":
		return fromDataAddresses(data.Addresses189, second)
	case "190":
		return fromDataAddresses(data.Addresses190, second)
	case "191":
		return fromDataAddresses(data.Addresses191, second)
	case "192":
		return fromDataAddresses(data.Addresses192, second)
	case "193":
		return fromDataAddresses(data.Addresses193, second)
	case "194":
		return fromDataAddresses(data.Addresses194, second)
	case "195":
		return fromDataAddresses(data.Addresses195, second)
	case "196":
		return fromDataAddresses(data.Addresses196, second)
	case "197":
		return fromDataAddresses(data.Addresses197, second)
	case "198":
		return fromDataAddresses(data.Addresses198, second)
	case "201":
		return fromDataAddresses(data.Addresses201, second)
	case "202":
		return fromDataAddresses(data.Addresses202, second)
	case "203":
		return fromDataAddresses(data.Addresses203, second)
	case "204":
		return fromDataAddresses(data.Addresses204, second)
	case "205":
		return fromDataAddresses(data.Addresses205, second)
	case "206":
		return fromDataAddresses(data.Addresses206, second)
	case "207":
		return fromDataAddresses(data.Addresses207, second)
	case "208":
		return fromDataAddresses(data.Addresses208, second)
	case "210":
		return fromDataAddresses(data.Addresses210, second)
	case "211":
		return fromDataAddresses(data.Addresses211, second)
	case "212":
		return fromDataAddresses(data.Addresses212, second)
	case "213":
		return fromDataAddresses(data.Addresses213, second)
	case "214":
		return fromDataAddresses(data.Addresses214, second)
	case "215":
		return fromDataAddresses(data.Addresses215, second)
	case "216":
		return fromDataAddresses(data.Addresses216, second)
	case "220":
		return fromDataAddresses(data.Addresses220, second)
	case "221":
		return fromDataAddresses(data.Addresses221, second)
	case "222":
		return fromDataAddresses(data.Addresses222, second)
	case "223":
		return fromDataAddresses(data.Addresses223, second)
	case "224":
		return fromDataAddresses(data.Addresses224, second)
	case "225":
		return fromDataAddresses(data.Addresses225, second)
	case "226":
		return fromDataAddresses(data.Addresses226, second)
	case "227":
		return fromDataAddresses(data.Addresses227, second)
	case "230":
		return fromDataAddresses(data.Addresses230, second)
	case "231":
		return fromDataAddresses(data.Addresses231, second)
	case "232":
		return fromDataAddresses(data.Addresses232, second)
	case "233":
		return fromDataAddresses(data.Addresses233, second)
	case "234":
		return fromDataAddresses(data.Addresses234, second)
	case "235":
		return fromDataAddresses(data.Addresses235, second)
	case "236":
		return fromDataAddresses(data.Addresses236, second)
	case "237":
		return fromDataAddresses(data.Addresses237, second)
	case "238":
		return fromDataAddresses(data.Addresses238, second)
	case "239":
		return fromDataAddresses(data.Addresses239, second)
	case "240":
		return fromDataAddresses(data.Addresses240, second)
	case "241":
		return fromDataAddresses(data.Addresses241, second)
	case "242":
		return fromDataAddresses(data.Addresses242, second)
	case "243":
		return fromDataAddresses(data.Addresses243, second)
	case "244":
		return fromDataAddresses(data.Addresses244, second)
	case "245":
		return fromDataAddresses(data.Addresses245, second)
	case "246":
		return fromDataAddresses(data.Addresses246, second)
	case "247":
		return fromDataAddresses(data.Addresses247, second)
	case "248":
		return fromDataAddresses(data.Addresses248, second)
	case "249":
		return fromDataAddresses(data.Addresses249, second)
	case "250":
		return fromDataAddresses(data.Addresses250, second)
	case "251":
		return fromDataAddresses(data.Addresses251, second)
	case "252":
		return fromDataAddresses(data.Addresses252, second)
	case "253":
		return fromDataAddresses(data.Addresses253, second)
	case "254":
		return fromDataAddresses(data.Addresses254, second)
	case "255":
		return fromDataAddresses(data.Addresses255, second)
	case "256":
		return fromDataAddresses(data.Addresses256, second)
	case "257":
		return fromDataAddresses(data.Addresses257, second)
	case "258":
		return fromDataAddresses(data.Addresses258, second)
	case "259":
		return fromDataAddresses(data.Addresses259, second)
	case "260":
		return fromDataAddresses(data.Addresses260, second)
	case "261":
		return fromDataAddresses(data.Addresses261, second)
	case "262":
		return fromDataAddresses(data.Addresses262, second)
	case "263":
		return fromDataAddresses(data.Addresses263, second)
	case "264":
		return fromDataAddresses(data.Addresses264, second)
	case "265":
		return fromDataAddresses(data.Addresses265, second)
	case "266":
		return fromDataAddresses(data.Addresses266, second)
	case "267":
		return fromDataAddresses(data.Addresses267, second)
	case "270":
		return fromDataAddresses(data.Addresses270, second)
	case "271":
		return fromDataAddresses(data.Addresses271, second)
	case "272":
		return fromDataAddresses(data.Addresses272, second)
	case "273":
		return fromDataAddresses(data.Addresses273, second)
	case "274":
		return fromDataAddresses(data.Addresses274, second)
	case "275":
		return fromDataAddresses(data.Addresses275, second)
	case "276":
		return fromDataAddresses(data.Addresses276, second)
	case "277":
		return fromDataAddresses(data.Addresses277, second)
	case "278":
		return fromDataAddresses(data.Addresses278, second)
	case "279":
		return fromDataAddresses(data.Addresses279, second)
	case "282":
		return fromDataAddresses(data.Addresses282, second)
	case "283":
		return fromDataAddresses(data.Addresses283, second)
	case "284":
		return fromDataAddresses(data.Addresses284, second)
	case "285":
		return fromDataAddresses(data.Addresses285, second)
	case "286":
		return fromDataAddresses(data.Addresses286, second)
	case "287":
		return fromDataAddresses(data.Addresses287, second)
	case "288":
		return fromDataAddresses(data.Addresses288, second)
	case "289":
		return fromDataAddresses(data.Addresses289, second)
	case "290":
		return fromDataAddresses(data.Addresses290, second)
	case "292":
		return fromDataAddresses(data.Addresses292, second)
	case "293":
		return fromDataAddresses(data.Addresses293, second)
	case "294":
		return fromDataAddresses(data.Addresses294, second)
	case "295":
		return fromDataAddresses(data.Addresses295, second)
	case "296":
		return fromDataAddresses(data.Addresses296, second)
	case "297":
		return fromDataAddresses(data.Addresses297, second)
	case "298":
		return fromDataAddresses(data.Addresses298, second)
	case "299":
		return fromDataAddresses(data.Addresses299, second)
	case "300":
		return fromDataAddresses(data.Addresses300, second)
	case "301":
		return fromDataAddresses(data.Addresses301, second)
	case "302":
		return fromDataAddresses(data.Addresses302, second)
	case "303":
		return fromDataAddresses(data.Addresses303, second)
	case "304":
		return fromDataAddresses(data.Addresses304, second)
	case "305":
		return fromDataAddresses(data.Addresses305, second)
	case "306":
		return fromDataAddresses(data.Addresses306, second)
	case "307":
		return fromDataAddresses(data.Addresses307, second)
	case "308":
		return fromDataAddresses(data.Addresses308, second)
	case "309":
		return fromDataAddresses(data.Addresses309, second)
	case "310":
		return fromDataAddresses(data.Addresses310, second)
	case "311":
		return fromDataAddresses(data.Addresses311, second)
	case "312":
		return fromDataAddresses(data.Addresses312, second)
	case "313":
		return fromDataAddresses(data.Addresses313, second)
	case "314":
		return fromDataAddresses(data.Addresses314, second)
	case "315":
		return fromDataAddresses(data.Addresses315, second)
	case "316":
		return fromDataAddresses(data.Addresses316, second)
	case "317":
		return fromDataAddresses(data.Addresses317, second)
	case "318":
		return fromDataAddresses(data.Addresses318, second)
	case "319":
		return fromDataAddresses(data.Addresses319, second)
	case "320":
		return fromDataAddresses(data.Addresses320, second)
	case "321":
		return fromDataAddresses(data.Addresses321, second)
	case "322":
		return fromDataAddresses(data.Addresses322, second)
	case "323":
		return fromDataAddresses(data.Addresses323, second)
	case "324":
		return fromDataAddresses(data.Addresses324, second)
	case "325":
		return fromDataAddresses(data.Addresses325, second)
	case "326":
		return fromDataAddresses(data.Addresses326, second)
	case "327":
		return fromDataAddresses(data.Addresses327, second)
	case "328":
		return fromDataAddresses(data.Addresses328, second)
	case "329":
		return fromDataAddresses(data.Addresses329, second)
	case "330":
		return fromDataAddresses(data.Addresses330, second)
	case "331":
		return fromDataAddresses(data.Addresses331, second)
	case "332":
		return fromDataAddresses(data.Addresses332, second)
	case "333":
		return fromDataAddresses(data.Addresses333, second)
	case "334":
		return fromDataAddresses(data.Addresses334, second)
	case "335":
		return fromDataAddresses(data.Addresses335, second)
	case "336":
		return fromDataAddresses(data.Addresses336, second)
	case "337":
		return fromDataAddresses(data.Addresses337, second)
	case "338":
		return fromDataAddresses(data.Addresses338, second)
	case "339":
		return fromDataAddresses(data.Addresses339, second)
	case "340":
		return fromDataAddresses(data.Addresses340, second)
	case "341":
		return fromDataAddresses(data.Addresses341, second)
	case "342":
		return fromDataAddresses(data.Addresses342, second)
	case "343":
		return fromDataAddresses(data.Addresses343, second)
	case "344":
		return fromDataAddresses(data.Addresses344, second)
	case "345":
		return fromDataAddresses(data.Addresses345, second)
	case "346":
		return fromDataAddresses(data.Addresses346, second)
	case "347":
		return fromDataAddresses(data.Addresses347, second)
	case "348":
		return fromDataAddresses(data.Addresses348, second)
	case "349":
		return fromDataAddresses(data.Addresses349, second)
	case "350":
		return fromDataAddresses(data.Addresses350, second)
	case "351":
		return fromDataAddresses(data.Addresses351, second)
	case "352":
		return fromDataAddresses(data.Addresses352, second)
	case "353":
		return fromDataAddresses(data.Addresses353, second)
	case "354":
		return fromDataAddresses(data.Addresses354, second)
	case "355":
		return fromDataAddresses(data.Addresses355, second)
	case "356":
		return fromDataAddresses(data.Addresses356, second)
	case "357":
		return fromDataAddresses(data.Addresses357, second)
	case "358":
		return fromDataAddresses(data.Addresses358, second)
	case "359":
		return fromDataAddresses(data.Addresses359, second)
	case "360":
		return fromDataAddresses(data.Addresses360, second)
	case "361":
		return fromDataAddresses(data.Addresses361, second)
	case "362":
		return fromDataAddresses(data.Addresses362, second)
	case "363":
		return fromDataAddresses(data.Addresses363, second)
	case "364":
		return fromDataAddresses(data.Addresses364, second)
	case "365":
		return fromDataAddresses(data.Addresses365, second)
	case "366":
		return fromDataAddresses(data.Addresses366, second)
	case "367":
		return fromDataAddresses(data.Addresses367, second)
	case "368":
		return fromDataAddresses(data.Addresses368, second)
	case "369":
		return fromDataAddresses(data.Addresses369, second)
	case "370":
		return fromDataAddresses(data.Addresses370, second)
	case "371":
		return fromDataAddresses(data.Addresses371, second)
	case "372":
		return fromDataAddresses(data.Addresses372, second)
	case "373":
		return fromDataAddresses(data.Addresses373, second)
	case "374":
		return fromDataAddresses(data.Addresses374, second)
	case "375":
		return fromDataAddresses(data.Addresses375, second)
	case "376":
		return fromDataAddresses(data.Addresses376, second)
	case "377":
		return fromDataAddresses(data.Addresses377, second)
	case "378":
		return fromDataAddresses(data.Addresses378, second)
	case "379":
		return fromDataAddresses(data.Addresses379, second)
	case "380":
		return fromDataAddresses(data.Addresses380, second)
	case "381":
		return fromDataAddresses(data.Addresses381, second)
	case "382":
		return fromDataAddresses(data.Addresses382, second)
	case "383":
		return fromDataAddresses(data.Addresses383, second)
	case "384":
		return fromDataAddresses(data.Addresses384, second)
	case "385":
		return fromDataAddresses(data.Addresses385, second)
	case "386":
		return fromDataAddresses(data.Addresses386, second)
	case "387":
		return fromDataAddresses(data.Addresses387, second)
	case "388":
		return fromDataAddresses(data.Addresses388, second)
	case "389":
		return fromDataAddresses(data.Addresses389, second)
	case "390":
		return fromDataAddresses(data.Addresses390, second)
	case "391":
		return fromDataAddresses(data.Addresses391, second)
	case "392":
		return fromDataAddresses(data.Addresses392, second)
	case "393":
		return fromDataAddresses(data.Addresses393, second)
	case "394":
		return fromDataAddresses(data.Addresses394, second)
	case "395":
		return fromDataAddresses(data.Addresses395, second)
	case "396":
		return fromDataAddresses(data.Addresses396, second)
	case "397":
		return fromDataAddresses(data.Addresses397, second)
	case "398":
		return fromDataAddresses(data.Addresses398, second)
	case "399":
		return fromDataAddresses(data.Addresses399, second)
	case "400":
		return fromDataAddresses(data.Addresses400, second)
	case "401":
		return fromDataAddresses(data.Addresses401, second)
	case "402":
		return fromDataAddresses(data.Addresses402, second)
	case "403":
		return fromDataAddresses(data.Addresses403, second)
	case "404":
		return fromDataAddresses(data.Addresses404, second)
	case "405":
		return fromDataAddresses(data.Addresses405, second)
	case "406":
		return fromDataAddresses(data.Addresses406, second)
	case "407":
		return fromDataAddresses(data.Addresses407, second)
	case "408":
		return fromDataAddresses(data.Addresses408, second)
	case "409":
		return fromDataAddresses(data.Addresses409, second)
	case "410":
		return fromDataAddresses(data.Addresses410, second)
	case "411":
		return fromDataAddresses(data.Addresses411, second)
	case "412":
		return fromDataAddresses(data.Addresses412, second)
	case "413":
		return fromDataAddresses(data.Addresses413, second)
	case "414":
		return fromDataAddresses(data.Addresses414, second)
	case "415":
		return fromDataAddresses(data.Addresses415, second)
	case "416":
		return fromDataAddresses(data.Addresses416, second)
	case "417":
		return fromDataAddresses(data.Addresses417, second)
	case "418":
		return fromDataAddresses(data.Addresses418, second)
	case "419":
		return fromDataAddresses(data.Addresses419, second)
	case "420":
		return fromDataAddresses(data.Addresses420, second)
	case "421":
		return fromDataAddresses(data.Addresses421, second)
	case "422":
		return fromDataAddresses(data.Addresses422, second)
	case "424":
		return fromDataAddresses(data.Addresses424, second)
	case "425":
		return fromDataAddresses(data.Addresses425, second)
	case "426":
		return fromDataAddresses(data.Addresses426, second)
	case "427":
		return fromDataAddresses(data.Addresses427, second)
	case "428":
		return fromDataAddresses(data.Addresses428, second)
	case "430":
		return fromDataAddresses(data.Addresses430, second)
	case "431":
		return fromDataAddresses(data.Addresses431, second)
	case "432":
		return fromDataAddresses(data.Addresses432, second)
	case "433":
		return fromDataAddresses(data.Addresses433, second)
	case "434":
		return fromDataAddresses(data.Addresses434, second)
	case "435":
		return fromDataAddresses(data.Addresses435, second)
	case "436":
		return fromDataAddresses(data.Addresses436, second)
	case "437":
		return fromDataAddresses(data.Addresses437, second)
	case "438":
		return fromDataAddresses(data.Addresses438, second)
	case "439":
		return fromDataAddresses(data.Addresses439, second)
	case "440":
		return fromDataAddresses(data.Addresses440, second)
	case "441":
		return fromDataAddresses(data.Addresses441, second)
	case "442":
		return fromDataAddresses(data.Addresses442, second)
	case "443":
		return fromDataAddresses(data.Addresses443, second)
	case "444":
		return fromDataAddresses(data.Addresses444, second)
	case "445":
		return fromDataAddresses(data.Addresses445, second)
	case "446":
		return fromDataAddresses(data.Addresses446, second)
	case "447":
		return fromDataAddresses(data.Addresses447, second)
	case "448":
		return fromDataAddresses(data.Addresses448, second)
	case "449":
		return fromDataAddresses(data.Addresses449, second)
	case "450":
		return fromDataAddresses(data.Addresses450, second)
	case "451":
		return fromDataAddresses(data.Addresses451, second)
	case "452":
		return fromDataAddresses(data.Addresses452, second)
	case "453":
		return fromDataAddresses(data.Addresses453, second)
	case "454":
		return fromDataAddresses(data.Addresses454, second)
	case "455":
		return fromDataAddresses(data.Addresses455, second)
	case "456":
		return fromDataAddresses(data.Addresses456, second)
	case "457":
		return fromDataAddresses(data.Addresses457, second)
	case "458":
		return fromDataAddresses(data.Addresses458, second)
	case "459":
		return fromDataAddresses(data.Addresses459, second)
	case "460":
		return fromDataAddresses(data.Addresses460, second)
	case "461":
		return fromDataAddresses(data.Addresses461, second)
	case "462":
		return fromDataAddresses(data.Addresses462, second)
	case "463":
		return fromDataAddresses(data.Addresses463, second)
	case "464":
		return fromDataAddresses(data.Addresses464, second)
	case "465":
		return fromDataAddresses(data.Addresses465, second)
	case "466":
		return fromDataAddresses(data.Addresses466, second)
	case "467":
		return fromDataAddresses(data.Addresses467, second)
	case "468":
		return fromDataAddresses(data.Addresses468, second)
	case "469":
		return fromDataAddresses(data.Addresses469, second)
	case "470":
		return fromDataAddresses(data.Addresses470, second)
	case "471":
		return fromDataAddresses(data.Addresses471, second)
	case "472":
		return fromDataAddresses(data.Addresses472, second)
	case "473":
		return fromDataAddresses(data.Addresses473, second)
	case "474":
		return fromDataAddresses(data.Addresses474, second)
	case "475":
		return fromDataAddresses(data.Addresses475, second)
	case "476":
		return fromDataAddresses(data.Addresses476, second)
	case "477":
		return fromDataAddresses(data.Addresses477, second)
	case "478":
		return fromDataAddresses(data.Addresses478, second)
	case "479":
		return fromDataAddresses(data.Addresses479, second)
	case "480":
		return fromDataAddresses(data.Addresses480, second)
	case "481":
		return fromDataAddresses(data.Addresses481, second)
	case "482":
		return fromDataAddresses(data.Addresses482, second)
	case "483":
		return fromDataAddresses(data.Addresses483, second)
	case "484":
		return fromDataAddresses(data.Addresses484, second)
	case "485":
		return fromDataAddresses(data.Addresses485, second)
	case "486":
		return fromDataAddresses(data.Addresses486, second)
	case "487":
		return fromDataAddresses(data.Addresses487, second)
	case "488":
		return fromDataAddresses(data.Addresses488, second)
	case "489":
		return fromDataAddresses(data.Addresses489, second)
	case "490":
		return fromDataAddresses(data.Addresses490, second)
	case "491":
		return fromDataAddresses(data.Addresses491, second)
	case "492":
		return fromDataAddresses(data.Addresses492, second)
	case "493":
		return fromDataAddresses(data.Addresses493, second)
	case "494":
		return fromDataAddresses(data.Addresses494, second)
	case "495":
		return fromDataAddresses(data.Addresses495, second)
	case "496":
		return fromDataAddresses(data.Addresses496, second)
	case "497":
		return fromDataAddresses(data.Addresses497, second)
	case "498":
		return fromDataAddresses(data.Addresses498, second)
	case "500":
		return fromDataAddresses(data.Addresses500, second)
	case "501":
		return fromDataAddresses(data.Addresses501, second)
	case "502":
		return fromDataAddresses(data.Addresses502, second)
	case "503":
		return fromDataAddresses(data.Addresses503, second)
	case "504":
		return fromDataAddresses(data.Addresses504, second)
	case "505":
		return fromDataAddresses(data.Addresses505, second)
	case "506":
		return fromDataAddresses(data.Addresses506, second)
	case "507":
		return fromDataAddresses(data.Addresses507, second)
	case "508":
		return fromDataAddresses(data.Addresses508, second)
	case "509":
		return fromDataAddresses(data.Addresses509, second)
	case "510":
		return fromDataAddresses(data.Addresses510, second)
	case "511":
		return fromDataAddresses(data.Addresses511, second)
	case "512":
		return fromDataAddresses(data.Addresses512, second)
	case "513":
		return fromDataAddresses(data.Addresses513, second)
	case "514":
		return fromDataAddresses(data.Addresses514, second)
	case "515":
		return fromDataAddresses(data.Addresses515, second)
	case "516":
		return fromDataAddresses(data.Addresses516, second)
	case "517":
		return fromDataAddresses(data.Addresses517, second)
	case "518":
		return fromDataAddresses(data.Addresses518, second)
	case "519":
		return fromDataAddresses(data.Addresses519, second)
	case "520":
		return fromDataAddresses(data.Addresses520, second)
	case "521":
		return fromDataAddresses(data.Addresses521, second)
	case "522":
		return fromDataAddresses(data.Addresses522, second)
	case "523":
		return fromDataAddresses(data.Addresses523, second)
	case "524":
		return fromDataAddresses(data.Addresses524, second)
	case "525":
		return fromDataAddresses(data.Addresses525, second)
	case "526":
		return fromDataAddresses(data.Addresses526, second)
	case "527":
		return fromDataAddresses(data.Addresses527, second)
	case "528":
		return fromDataAddresses(data.Addresses528, second)
	case "529":
		return fromDataAddresses(data.Addresses529, second)
	case "530":
		return fromDataAddresses(data.Addresses530, second)
	case "531":
		return fromDataAddresses(data.Addresses531, second)
	case "532":
		return fromDataAddresses(data.Addresses532, second)
	case "533":
		return fromDataAddresses(data.Addresses533, second)
	case "534":
		return fromDataAddresses(data.Addresses534, second)
	case "535":
		return fromDataAddresses(data.Addresses535, second)
	case "536":
		return fromDataAddresses(data.Addresses536, second)
	case "537":
		return fromDataAddresses(data.Addresses537, second)
	case "538":
		return fromDataAddresses(data.Addresses538, second)
	case "539":
		return fromDataAddresses(data.Addresses539, second)
	case "540":
		return fromDataAddresses(data.Addresses540, second)
	case "541":
		return fromDataAddresses(data.Addresses541, second)
	case "542":
		return fromDataAddresses(data.Addresses542, second)
	case "543":
		return fromDataAddresses(data.Addresses543, second)
	case "544":
		return fromDataAddresses(data.Addresses544, second)
	case "545":
		return fromDataAddresses(data.Addresses545, second)
	case "546":
		return fromDataAddresses(data.Addresses546, second)
	case "547":
		return fromDataAddresses(data.Addresses547, second)
	case "549":
		return fromDataAddresses(data.Addresses549, second)
	case "550":
		return fromDataAddresses(data.Addresses550, second)
	case "551":
		return fromDataAddresses(data.Addresses551, second)
	case "552":
		return fromDataAddresses(data.Addresses552, second)
	case "553":
		return fromDataAddresses(data.Addresses553, second)
	case "554":
		return fromDataAddresses(data.Addresses554, second)
	case "555":
		return fromDataAddresses(data.Addresses555, second)
	case "556":
		return fromDataAddresses(data.Addresses556, second)
	case "557":
		return fromDataAddresses(data.Addresses557, second)
	case "558":
		return fromDataAddresses(data.Addresses558, second)
	case "559":
		return fromDataAddresses(data.Addresses559, second)
	case "560":
		return fromDataAddresses(data.Addresses560, second)
	case "561":
		return fromDataAddresses(data.Addresses561, second)
	case "562":
		return fromDataAddresses(data.Addresses562, second)
	case "563":
		return fromDataAddresses(data.Addresses563, second)
	case "564":
		return fromDataAddresses(data.Addresses564, second)
	case "565":
		return fromDataAddresses(data.Addresses565, second)
	case "566":
		return fromDataAddresses(data.Addresses566, second)
	case "567":
		return fromDataAddresses(data.Addresses567, second)
	case "568":
		return fromDataAddresses(data.Addresses568, second)
	case "569":
		return fromDataAddresses(data.Addresses569, second)
	case "570":
		return fromDataAddresses(data.Addresses570, second)
	case "571":
		return fromDataAddresses(data.Addresses571, second)
	case "572":
		return fromDataAddresses(data.Addresses572, second)
	case "573":
		return fromDataAddresses(data.Addresses573, second)
	case "574":
		return fromDataAddresses(data.Addresses574, second)
	case "575":
		return fromDataAddresses(data.Addresses575, second)
	case "576":
		return fromDataAddresses(data.Addresses576, second)
	case "577":
		return fromDataAddresses(data.Addresses577, second)
	case "578":
		return fromDataAddresses(data.Addresses578, second)
	case "579":
		return fromDataAddresses(data.Addresses579, second)
	case "580":
		return fromDataAddresses(data.Addresses580, second)
	case "581":
		return fromDataAddresses(data.Addresses581, second)
	case "582":
		return fromDataAddresses(data.Addresses582, second)
	case "583":
		return fromDataAddresses(data.Addresses583, second)
	case "584":
		return fromDataAddresses(data.Addresses584, second)
	case "585":
		return fromDataAddresses(data.Addresses585, second)
	case "586":
		return fromDataAddresses(data.Addresses586, second)
	case "587":
		return fromDataAddresses(data.Addresses587, second)
	case "589":
		return fromDataAddresses(data.Addresses589, second)
	case "590":
		return fromDataAddresses(data.Addresses590, second)
	case "591":
		return fromDataAddresses(data.Addresses591, second)
	case "592":
		return fromDataAddresses(data.Addresses592, second)
	case "593":
		return fromDataAddresses(data.Addresses593, second)
	case "594":
		return fromDataAddresses(data.Addresses594, second)
	case "595":
		return fromDataAddresses(data.Addresses595, second)
	case "596":
		return fromDataAddresses(data.Addresses596, second)
	case "597":
		return fromDataAddresses(data.Addresses597, second)
	case "598":
		return fromDataAddresses(data.Addresses598, second)
	case "599":
		return fromDataAddresses(data.Addresses599, second)
	case "600":
		return fromDataAddresses(data.Addresses600, second)
	case "601":
		return fromDataAddresses(data.Addresses601, second)
	case "602":
		return fromDataAddresses(data.Addresses602, second)
	case "603":
		return fromDataAddresses(data.Addresses603, second)
	case "604":
		return fromDataAddresses(data.Addresses604, second)
	case "605":
		return fromDataAddresses(data.Addresses605, second)
	case "606":
		return fromDataAddresses(data.Addresses606, second)
	case "607":
		return fromDataAddresses(data.Addresses607, second)
	case "610":
		return fromDataAddresses(data.Addresses610, second)
	case "611":
		return fromDataAddresses(data.Addresses611, second)
	case "612":
		return fromDataAddresses(data.Addresses612, second)
	case "613":
		return fromDataAddresses(data.Addresses613, second)
	case "614":
		return fromDataAddresses(data.Addresses614, second)
	case "615":
		return fromDataAddresses(data.Addresses615, second)
	case "616":
		return fromDataAddresses(data.Addresses616, second)
	case "617":
		return fromDataAddresses(data.Addresses617, second)
	case "618":
		return fromDataAddresses(data.Addresses618, second)
	case "619":
		return fromDataAddresses(data.Addresses619, second)
	case "620":
		return fromDataAddresses(data.Addresses620, second)
	case "621":
		return fromDataAddresses(data.Addresses621, second)
	case "622":
		return fromDataAddresses(data.Addresses622, second)
	case "623":
		return fromDataAddresses(data.Addresses623, second)
	case "624":
		return fromDataAddresses(data.Addresses624, second)
	case "625":
		return fromDataAddresses(data.Addresses625, second)
	case "626":
		return fromDataAddresses(data.Addresses626, second)
	case "627":
		return fromDataAddresses(data.Addresses627, second)
	case "629":
		return fromDataAddresses(data.Addresses629, second)
	case "630":
		return fromDataAddresses(data.Addresses630, second)
	case "631":
		return fromDataAddresses(data.Addresses631, second)
	case "632":
		return fromDataAddresses(data.Addresses632, second)
	case "633":
		return fromDataAddresses(data.Addresses633, second)
	case "634":
		return fromDataAddresses(data.Addresses634, second)
	case "635":
		return fromDataAddresses(data.Addresses635, second)
	case "636":
		return fromDataAddresses(data.Addresses636, second)
	case "637":
		return fromDataAddresses(data.Addresses637, second)
	case "638":
		return fromDataAddresses(data.Addresses638, second)
	case "639":
		return fromDataAddresses(data.Addresses639, second)
	case "640":
		return fromDataAddresses(data.Addresses640, second)
	case "641":
		return fromDataAddresses(data.Addresses641, second)
	case "642":
		return fromDataAddresses(data.Addresses642, second)
	case "643":
		return fromDataAddresses(data.Addresses643, second)
	case "644":
		return fromDataAddresses(data.Addresses644, second)
	case "645":
		return fromDataAddresses(data.Addresses645, second)
	case "646":
		return fromDataAddresses(data.Addresses646, second)
	case "647":
		return fromDataAddresses(data.Addresses647, second)
	case "648":
		return fromDataAddresses(data.Addresses648, second)
	case "649":
		return fromDataAddresses(data.Addresses649, second)
	case "650":
		return fromDataAddresses(data.Addresses650, second)
	case "651":
		return fromDataAddresses(data.Addresses651, second)
	case "652":
		return fromDataAddresses(data.Addresses652, second)
	case "653":
		return fromDataAddresses(data.Addresses653, second)
	case "654":
		return fromDataAddresses(data.Addresses654, second)
	case "655":
		return fromDataAddresses(data.Addresses655, second)
	case "656":
		return fromDataAddresses(data.Addresses656, second)
	case "657":
		return fromDataAddresses(data.Addresses657, second)
	case "658":
		return fromDataAddresses(data.Addresses658, second)
	case "659":
		return fromDataAddresses(data.Addresses659, second)
	case "660":
		return fromDataAddresses(data.Addresses660, second)
	case "661":
		return fromDataAddresses(data.Addresses661, second)
	case "662":
		return fromDataAddresses(data.Addresses662, second)
	case "663":
		return fromDataAddresses(data.Addresses663, second)
	case "664":
		return fromDataAddresses(data.Addresses664, second)
	case "665":
		return fromDataAddresses(data.Addresses665, second)
	case "666":
		return fromDataAddresses(data.Addresses666, second)
	case "667":
		return fromDataAddresses(data.Addresses667, second)
	case "668":
		return fromDataAddresses(data.Addresses668, second)
	case "669":
		return fromDataAddresses(data.Addresses669, second)
	case "670":
		return fromDataAddresses(data.Addresses670, second)
	case "671":
		return fromDataAddresses(data.Addresses671, second)
	case "672":
		return fromDataAddresses(data.Addresses672, second)
	case "673":
		return fromDataAddresses(data.Addresses673, second)
	case "674":
		return fromDataAddresses(data.Addresses674, second)
	case "675":
		return fromDataAddresses(data.Addresses675, second)
	case "676":
		return fromDataAddresses(data.Addresses676, second)
	case "677":
		return fromDataAddresses(data.Addresses677, second)
	case "678":
		return fromDataAddresses(data.Addresses678, second)
	case "679":
		return fromDataAddresses(data.Addresses679, second)
	case "680":
		return fromDataAddresses(data.Addresses680, second)
	case "681":
		return fromDataAddresses(data.Addresses681, second)
	case "682":
		return fromDataAddresses(data.Addresses682, second)
	case "683":
		return fromDataAddresses(data.Addresses683, second)
	case "684":
		return fromDataAddresses(data.Addresses684, second)
	case "685":
		return fromDataAddresses(data.Addresses685, second)
	case "689":
		return fromDataAddresses(data.Addresses689, second)
	case "690":
		return fromDataAddresses(data.Addresses690, second)
	case "691":
		return fromDataAddresses(data.Addresses691, second)
	case "692":
		return fromDataAddresses(data.Addresses692, second)
	case "693":
		return fromDataAddresses(data.Addresses693, second)
	case "694":
		return fromDataAddresses(data.Addresses694, second)
	case "695":
		return fromDataAddresses(data.Addresses695, second)
	case "696":
		return fromDataAddresses(data.Addresses696, second)
	case "697":
		return fromDataAddresses(data.Addresses697, second)
	case "698":
		return fromDataAddresses(data.Addresses698, second)
	case "699":
		return fromDataAddresses(data.Addresses699, second)
	case "700":
		return fromDataAddresses(data.Addresses700, second)
	case "701":
		return fromDataAddresses(data.Addresses701, second)
	case "702":
		return fromDataAddresses(data.Addresses702, second)
	case "703":
		return fromDataAddresses(data.Addresses703, second)
	case "704":
		return fromDataAddresses(data.Addresses704, second)
	case "705":
		return fromDataAddresses(data.Addresses705, second)
	case "706":
		return fromDataAddresses(data.Addresses706, second)
	case "707":
		return fromDataAddresses(data.Addresses707, second)
	case "708":
		return fromDataAddresses(data.Addresses708, second)
	case "709":
		return fromDataAddresses(data.Addresses709, second)
	case "710":
		return fromDataAddresses(data.Addresses710, second)
	case "711":
		return fromDataAddresses(data.Addresses711, second)
	case "712":
		return fromDataAddresses(data.Addresses712, second)
	case "713":
		return fromDataAddresses(data.Addresses713, second)
	case "714":
		return fromDataAddresses(data.Addresses714, second)
	case "715":
		return fromDataAddresses(data.Addresses715, second)
	case "716":
		return fromDataAddresses(data.Addresses716, second)
	case "717":
		return fromDataAddresses(data.Addresses717, second)
	case "718":
		return fromDataAddresses(data.Addresses718, second)
	case "719":
		return fromDataAddresses(data.Addresses719, second)
	case "720":
		return fromDataAddresses(data.Addresses720, second)
	case "721":
		return fromDataAddresses(data.Addresses721, second)
	case "722":
		return fromDataAddresses(data.Addresses722, second)
	case "723":
		return fromDataAddresses(data.Addresses723, second)
	case "725":
		return fromDataAddresses(data.Addresses725, second)
	case "726":
		return fromDataAddresses(data.Addresses726, second)
	case "727":
		return fromDataAddresses(data.Addresses727, second)
	case "728":
		return fromDataAddresses(data.Addresses728, second)
	case "729":
		return fromDataAddresses(data.Addresses729, second)
	case "730":
		return fromDataAddresses(data.Addresses730, second)
	case "731":
		return fromDataAddresses(data.Addresses731, second)
	case "732":
		return fromDataAddresses(data.Addresses732, second)
	case "733":
		return fromDataAddresses(data.Addresses733, second)
	case "734":
		return fromDataAddresses(data.Addresses734, second)
	case "735":
		return fromDataAddresses(data.Addresses735, second)
	case "736":
		return fromDataAddresses(data.Addresses736, second)
	case "737":
		return fromDataAddresses(data.Addresses737, second)
	case "738":
		return fromDataAddresses(data.Addresses738, second)
	case "739":
		return fromDataAddresses(data.Addresses739, second)
	case "740":
		return fromDataAddresses(data.Addresses740, second)
	case "741":
		return fromDataAddresses(data.Addresses741, second)
	case "742":
		return fromDataAddresses(data.Addresses742, second)
	case "743":
		return fromDataAddresses(data.Addresses743, second)
	case "744":
		return fromDataAddresses(data.Addresses744, second)
	case "745":
		return fromDataAddresses(data.Addresses745, second)
	case "746":
		return fromDataAddresses(data.Addresses746, second)
	case "747":
		return fromDataAddresses(data.Addresses747, second)
	case "749":
		return fromDataAddresses(data.Addresses749, second)
	case "750":
		return fromDataAddresses(data.Addresses750, second)
	case "751":
		return fromDataAddresses(data.Addresses751, second)
	case "752":
		return fromDataAddresses(data.Addresses752, second)
	case "753":
		return fromDataAddresses(data.Addresses753, second)
	case "754":
		return fromDataAddresses(data.Addresses754, second)
	case "755":
		return fromDataAddresses(data.Addresses755, second)
	case "756":
		return fromDataAddresses(data.Addresses756, second)
	case "757":
		return fromDataAddresses(data.Addresses757, second)
	case "758":
		return fromDataAddresses(data.Addresses758, second)
	case "759":
		return fromDataAddresses(data.Addresses759, second)
	case "760":
		return fromDataAddresses(data.Addresses760, second)
	case "761":
		return fromDataAddresses(data.Addresses761, second)
	case "762":
		return fromDataAddresses(data.Addresses762, second)
	case "763":
		return fromDataAddresses(data.Addresses763, second)
	case "764":
		return fromDataAddresses(data.Addresses764, second)
	case "765":
		return fromDataAddresses(data.Addresses765, second)
	case "766":
		return fromDataAddresses(data.Addresses766, second)
	case "767":
		return fromDataAddresses(data.Addresses767, second)
	case "768":
		return fromDataAddresses(data.Addresses768, second)
	case "769":
		return fromDataAddresses(data.Addresses769, second)
	case "770":
		return fromDataAddresses(data.Addresses770, second)
	case "771":
		return fromDataAddresses(data.Addresses771, second)
	case "772":
		return fromDataAddresses(data.Addresses772, second)
	case "773":
		return fromDataAddresses(data.Addresses773, second)
	case "774":
		return fromDataAddresses(data.Addresses774, second)
	case "775":
		return fromDataAddresses(data.Addresses775, second)
	case "776":
		return fromDataAddresses(data.Addresses776, second)
	case "777":
		return fromDataAddresses(data.Addresses777, second)
	case "778":
		return fromDataAddresses(data.Addresses778, second)
	case "779":
		return fromDataAddresses(data.Addresses779, second)
	case "780":
		return fromDataAddresses(data.Addresses780, second)
	case "781":
		return fromDataAddresses(data.Addresses781, second)
	case "782":
		return fromDataAddresses(data.Addresses782, second)
	case "783":
		return fromDataAddresses(data.Addresses783, second)
	case "784":
		return fromDataAddresses(data.Addresses784, second)
	case "785":
		return fromDataAddresses(data.Addresses785, second)
	case "786":
		return fromDataAddresses(data.Addresses786, second)
	case "787":
		return fromDataAddresses(data.Addresses787, second)
	case "788":
		return fromDataAddresses(data.Addresses788, second)
	case "789":
		return fromDataAddresses(data.Addresses789, second)
	case "790":
		return fromDataAddresses(data.Addresses790, second)
	case "791":
		return fromDataAddresses(data.Addresses791, second)
	case "792":
		return fromDataAddresses(data.Addresses792, second)
	case "793":
		return fromDataAddresses(data.Addresses793, second)
	case "794":
		return fromDataAddresses(data.Addresses794, second)
	case "795":
		return fromDataAddresses(data.Addresses795, second)
	case "796":
		return fromDataAddresses(data.Addresses796, second)
	case "797":
		return fromDataAddresses(data.Addresses797, second)
	case "798":
		return fromDataAddresses(data.Addresses798, second)
	case "799":
		return fromDataAddresses(data.Addresses799, second)
	case "800":
		return fromDataAddresses(data.Addresses800, second)
	case "801":
		return fromDataAddresses(data.Addresses801, second)
	case "802":
		return fromDataAddresses(data.Addresses802, second)
	case "803":
		return fromDataAddresses(data.Addresses803, second)
	case "804":
		return fromDataAddresses(data.Addresses804, second)
	case "805":
		return fromDataAddresses(data.Addresses805, second)
	case "806":
		return fromDataAddresses(data.Addresses806, second)
	case "807":
		return fromDataAddresses(data.Addresses807, second)
	case "808":
		return fromDataAddresses(data.Addresses808, second)
	case "809":
		return fromDataAddresses(data.Addresses809, second)
	case "810":
		return fromDataAddresses(data.Addresses810, second)
	case "811":
		return fromDataAddresses(data.Addresses811, second)
	case "812":
		return fromDataAddresses(data.Addresses812, second)
	case "813":
		return fromDataAddresses(data.Addresses813, second)
	case "814":
		return fromDataAddresses(data.Addresses814, second)
	case "815":
		return fromDataAddresses(data.Addresses815, second)
	case "816":
		return fromDataAddresses(data.Addresses816, second)
	case "817":
		return fromDataAddresses(data.Addresses817, second)
	case "818":
		return fromDataAddresses(data.Addresses818, second)
	case "819":
		return fromDataAddresses(data.Addresses819, second)
	case "820":
		return fromDataAddresses(data.Addresses820, second)
	case "821":
		return fromDataAddresses(data.Addresses821, second)
	case "822":
		return fromDataAddresses(data.Addresses822, second)
	case "823":
		return fromDataAddresses(data.Addresses823, second)
	case "824":
		return fromDataAddresses(data.Addresses824, second)
	case "825":
		return fromDataAddresses(data.Addresses825, second)
	case "826":
		return fromDataAddresses(data.Addresses826, second)
	case "827":
		return fromDataAddresses(data.Addresses827, second)
	case "828":
		return fromDataAddresses(data.Addresses828, second)
	case "829":
		return fromDataAddresses(data.Addresses829, second)
	case "830":
		return fromDataAddresses(data.Addresses830, second)
	case "831":
		return fromDataAddresses(data.Addresses831, second)
	case "832":
		return fromDataAddresses(data.Addresses832, second)
	case "833":
		return fromDataAddresses(data.Addresses833, second)
	case "834":
		return fromDataAddresses(data.Addresses834, second)
	case "835":
		return fromDataAddresses(data.Addresses835, second)
	case "836":
		return fromDataAddresses(data.Addresses836, second)
	case "837":
		return fromDataAddresses(data.Addresses837, second)
	case "838":
		return fromDataAddresses(data.Addresses838, second)
	case "839":
		return fromDataAddresses(data.Addresses839, second)
	case "840":
		return fromDataAddresses(data.Addresses840, second)
	case "841":
		return fromDataAddresses(data.Addresses841, second)
	case "842":
		return fromDataAddresses(data.Addresses842, second)
	case "843":
		return fromDataAddresses(data.Addresses843, second)
	case "844":
		return fromDataAddresses(data.Addresses844, second)
	case "845":
		return fromDataAddresses(data.Addresses845, second)
	case "846":
		return fromDataAddresses(data.Addresses846, second)
	case "847":
		return fromDataAddresses(data.Addresses847, second)
	case "848":
		return fromDataAddresses(data.Addresses848, second)
	case "849":
		return fromDataAddresses(data.Addresses849, second)
	case "850":
		return fromDataAddresses(data.Addresses850, second)
	case "851":
		return fromDataAddresses(data.Addresses851, second)
	case "852":
		return fromDataAddresses(data.Addresses852, second)
	case "853":
		return fromDataAddresses(data.Addresses853, second)
	case "854":
		return fromDataAddresses(data.Addresses854, second)
	case "855":
		return fromDataAddresses(data.Addresses855, second)
	case "856":
		return fromDataAddresses(data.Addresses856, second)
	case "857":
		return fromDataAddresses(data.Addresses857, second)
	case "858":
		return fromDataAddresses(data.Addresses858, second)
	case "859":
		return fromDataAddresses(data.Addresses859, second)
	case "860":
		return fromDataAddresses(data.Addresses860, second)
	case "861":
		return fromDataAddresses(data.Addresses861, second)
	case "862":
		return fromDataAddresses(data.Addresses862, second)
	case "863":
		return fromDataAddresses(data.Addresses863, second)
	case "864":
		return fromDataAddresses(data.Addresses864, second)
	case "865":
		return fromDataAddresses(data.Addresses865, second)
	case "866":
		return fromDataAddresses(data.Addresses866, second)
	case "867":
		return fromDataAddresses(data.Addresses867, second)
	case "868":
		return fromDataAddresses(data.Addresses868, second)
	case "869":
		return fromDataAddresses(data.Addresses869, second)
	case "870":
		return fromDataAddresses(data.Addresses870, second)
	case "871":
		return fromDataAddresses(data.Addresses871, second)
	case "872":
		return fromDataAddresses(data.Addresses872, second)
	case "873":
		return fromDataAddresses(data.Addresses873, second)
	case "874":
		return fromDataAddresses(data.Addresses874, second)
	case "875":
		return fromDataAddresses(data.Addresses875, second)
	case "876":
		return fromDataAddresses(data.Addresses876, second)
	case "877":
		return fromDataAddresses(data.Addresses877, second)
	case "878":
		return fromDataAddresses(data.Addresses878, second)
	case "879":
		return fromDataAddresses(data.Addresses879, second)
	case "880":
		return fromDataAddresses(data.Addresses880, second)
	case "881":
		return fromDataAddresses(data.Addresses881, second)
	case "882":
		return fromDataAddresses(data.Addresses882, second)
	case "883":
		return fromDataAddresses(data.Addresses883, second)
	case "884":
		return fromDataAddresses(data.Addresses884, second)
	case "885":
		return fromDataAddresses(data.Addresses885, second)
	case "886":
		return fromDataAddresses(data.Addresses886, second)
	case "887":
		return fromDataAddresses(data.Addresses887, second)
	case "888":
		return fromDataAddresses(data.Addresses888, second)
	case "889":
		return fromDataAddresses(data.Addresses889, second)
	case "890":
		return fromDataAddresses(data.Addresses890, second)
	case "891":
		return fromDataAddresses(data.Addresses891, second)
	case "892":
		return fromDataAddresses(data.Addresses892, second)
	case "893":
		return fromDataAddresses(data.Addresses893, second)
	case "894":
		return fromDataAddresses(data.Addresses894, second)
	case "895":
		return fromDataAddresses(data.Addresses895, second)
	case "896":
		return fromDataAddresses(data.Addresses896, second)
	case "897":
		return fromDataAddresses(data.Addresses897, second)
	case "898":
		return fromDataAddresses(data.Addresses898, second)
	case "899":
		return fromDataAddresses(data.Addresses899, second)
	case "900":
		return fromDataAddresses(data.Addresses900, second)
	case "901":
		return fromDataAddresses(data.Addresses901, second)
	case "902":
		return fromDataAddresses(data.Addresses902, second)
	case "903":
		return fromDataAddresses(data.Addresses903, second)
	case "904":
		return fromDataAddresses(data.Addresses904, second)
	case "905":
		return fromDataAddresses(data.Addresses905, second)
	case "906":
		return fromDataAddresses(data.Addresses906, second)
	case "907":
		return fromDataAddresses(data.Addresses907, second)
	case "910":
		return fromDataAddresses(data.Addresses910, second)
	case "911":
		return fromDataAddresses(data.Addresses911, second)
	case "912":
		return fromDataAddresses(data.Addresses912, second)
	case "913":
		return fromDataAddresses(data.Addresses913, second)
	case "914":
		return fromDataAddresses(data.Addresses914, second)
	case "915":
		return fromDataAddresses(data.Addresses915, second)
	case "916":
		return fromDataAddresses(data.Addresses916, second)
	case "917":
		return fromDataAddresses(data.Addresses917, second)
	case "918":
		return fromDataAddresses(data.Addresses918, second)
	case "919":
		return fromDataAddresses(data.Addresses919, second)
	case "920":
		return fromDataAddresses(data.Addresses920, second)
	case "921":
		return fromDataAddresses(data.Addresses921, second)
	case "922":
		return fromDataAddresses(data.Addresses922, second)
	case "923":
		return fromDataAddresses(data.Addresses923, second)
	case "924":
		return fromDataAddresses(data.Addresses924, second)
	case "925":
		return fromDataAddresses(data.Addresses925, second)
	case "926":
		return fromDataAddresses(data.Addresses926, second)
	case "927":
		return fromDataAddresses(data.Addresses927, second)
	case "928":
		return fromDataAddresses(data.Addresses928, second)
	case "929":
		return fromDataAddresses(data.Addresses929, second)
	case "930":
		return fromDataAddresses(data.Addresses930, second)
	case "931":
		return fromDataAddresses(data.Addresses931, second)
	case "932":
		return fromDataAddresses(data.Addresses932, second)
	case "933":
		return fromDataAddresses(data.Addresses933, second)
	case "934":
		return fromDataAddresses(data.Addresses934, second)
	case "935":
		return fromDataAddresses(data.Addresses935, second)
	case "936":
		return fromDataAddresses(data.Addresses936, second)
	case "937":
		return fromDataAddresses(data.Addresses937, second)
	case "938":
		return fromDataAddresses(data.Addresses938, second)
	case "939":
		return fromDataAddresses(data.Addresses939, second)
	case "940":
		return fromDataAddresses(data.Addresses940, second)
	case "941":
		return fromDataAddresses(data.Addresses941, second)
	case "942":
		return fromDataAddresses(data.Addresses942, second)
	case "943":
		return fromDataAddresses(data.Addresses943, second)
	case "944":
		return fromDataAddresses(data.Addresses944, second)
	case "945":
		return fromDataAddresses(data.Addresses945, second)
	case "946":
		return fromDataAddresses(data.Addresses946, second)
	case "947":
		return fromDataAddresses(data.Addresses947, second)
	case "948":
		return fromDataAddresses(data.Addresses948, second)
	case "949":
		return fromDataAddresses(data.Addresses949, second)
	case "950":
		return fromDataAddresses(data.Addresses950, second)
	case "951":
		return fromDataAddresses(data.Addresses951, second)
	case "952":
		return fromDataAddresses(data.Addresses952, second)
	case "953":
		return fromDataAddresses(data.Addresses953, second)
	case "954":
		return fromDataAddresses(data.Addresses954, second)
	case "955":
		return fromDataAddresses(data.Addresses955, second)
	case "956":
		return fromDataAddresses(data.Addresses956, second)
	case "957":
		return fromDataAddresses(data.Addresses957, second)
	case "958":
		return fromDataAddresses(data.Addresses958, second)
	case "959":
		return fromDataAddresses(data.Addresses959, second)
	case "960":
		return fromDataAddresses(data.Addresses960, second)
	case "961":
		return fromDataAddresses(data.Addresses961, second)
	case "962":
		return fromDataAddresses(data.Addresses962, second)
	case "963":
		return fromDataAddresses(data.Addresses963, second)
	case "964":
		return fromDataAddresses(data.Addresses964, second)
	case "965":
		return fromDataAddresses(data.Addresses965, second)
	case "966":
		return fromDataAddresses(data.Addresses966, second)
	case "967":
		return fromDataAddresses(data.Addresses967, second)
	case "968":
		return fromDataAddresses(data.Addresses968, second)
	case "969":
		return fromDataAddresses(data.Addresses969, second)
	case "970":
		return fromDataAddresses(data.Addresses970, second)
	case "971":
		return fromDataAddresses(data.Addresses971, second)
	case "972":
		return fromDataAddresses(data.Addresses972, second)
	case "973":
		return fromDataAddresses(data.Addresses973, second)
	case "974":
		return fromDataAddresses(data.Addresses974, second)
	case "975":
		return fromDataAddresses(data.Addresses975, second)
	case "976":
		return fromDataAddresses(data.Addresses976, second)
	case "979":
		return fromDataAddresses(data.Addresses979, second)
	case "980":
		return fromDataAddresses(data.Addresses980, second)
	case "981":
		return fromDataAddresses(data.Addresses981, second)
	case "982":
		return fromDataAddresses(data.Addresses982, second)
	case "983":
		return fromDataAddresses(data.Addresses983, second)
	case "984":
		return fromDataAddresses(data.Addresses984, second)
	case "985":
		return fromDataAddresses(data.Addresses985, second)
	case "986":
		return fromDataAddresses(data.Addresses986, second)
	case "987":
		return fromDataAddresses(data.Addresses987, second)
	case "988":
		return fromDataAddresses(data.Addresses988, second)
	case "989":
		return fromDataAddresses(data.Addresses989, second)
	case "990":
		return fromDataAddresses(data.Addresses990, second)
	case "991":
		return fromDataAddresses(data.Addresses991, second)
	case "992":
		return fromDataAddresses(data.Addresses992, second)
	case "993":
		return fromDataAddresses(data.Addresses993, second)
	case "994":
		return fromDataAddresses(data.Addresses994, second)
	case "995":
		return fromDataAddresses(data.Addresses995, second)
	case "996":
		return fromDataAddresses(data.Addresses996, second)
	case "997":
		return fromDataAddresses(data.Addresses997, second)
	case "998":
		return fromDataAddresses(data.Addresses998, second)
	case "999":
		return fromDataAddresses(data.Addresses999, second)
	}
	return nil, ErrNotFound
}

func fromDataAddresses(m map[string][]*address.Address, p string) ([]*Address, error) {
	d, ok := m[p]
	if !ok {
		return nil, ErrNotFound
	}
	r := make([]*Address, len(d))
	for i, a := range d {
		r[i] = (*Address)(a)
	}
	return r, nil
}
