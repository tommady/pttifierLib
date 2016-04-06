package pttifierLib_test

var TestBoardHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
<meta name="viewport" content="width=device-width">
<title>看板 LoL 文章列表 - 批踢踢實業坊</title>
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs-common.css">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs.css" media="screen">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/pushstream.css" media="screen">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs-print.css" media="print">
<script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
<script src="//images.ptt.cc/v2.14/bbs.js"></script>
<script type="text/javascript">
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-32365737-1']);
  _gaq.push(['_setDomainName', 'ptt.cc']);
  _gaq.push(['_trackPageview']);
  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
</script>
	</head>
    <body>
<div id="topbar-container">
	<div id="topbar" class="bbs-content">
		<a id="logo" href="/">批踢踢實業坊</a>
		<span>&rsaquo;</span>
		<a class="board" href="/bbs/LoL/index.html"><span class="board-label">看板 </span>LoL</a>
		<a class="right small" href="/about.html">關於我們</a>
		<a class="right small" href="/contact.html">聯絡資訊</a>
	</div>
</div>
<div id="main-container">
	<div id="action-bar-container">
		<div class="action-bar">
			<div class="btn-group">
				<a class="btn selected" href="/bbs/LoL/index.html">看板</a>
				<a class="btn" href="/man/LoL/index.html">精華區</a>
			</div>
			<div class="btn-group pull-right">
				<a class="btn wide" href="/bbs/LoL/index1.html">最舊</a>
				<a class="btn wide" href="/bbs/LoL/index4380.html">&lsaquo; 上頁</a>
				<a class="btn wide" href="/bbs/LoL/index4382.html">下頁 &rsaquo;</a>
				<a class="btn wide" href="/bbs/LoL/index.html">最新</a>
			</div>
		</div>
	</div>
	<div class="r-list-container bbs-screen">
		<div class="r-ent">
			<div class="nrec"><span class="hl f3">11</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1457622048.A.D80.html">[公告] LoL 樂透開獎</a>
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">[彩券]</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f3">11</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1457622048.A.Z80.html">[公告] LoL 小樂透開獎</a>
			</div>
			<div class="meta">
				<div class="date"> 3/11</div>
				<div class="author">[彩券]</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f2">4</span></div>
			<div class="mark"></div>
			<div class="title">
				<a href="/bbs/LoL/M.1457622282.A.D4F.html">Re: [問題] AlphaGo打LoL的話會怎麼樣</a>
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">lzainside</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"></div>
			<div class="mark"></div>
			<div class="title">
				(本文已被刪除) [asdxxg5]
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">-</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f2">2</span></div>
			<div class="mark"></div>
			<div class="title">
				<a href="/bbs/LoL/M.1457622407.A.C98.html">[問題] 關於四名中路的比較！</a>
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">lizst980074</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f2">1</span></div>
			<div class="mark"></div>
			<div class="title">
				<a href="/bbs/LoL/M.1457622551.A.650.html">[閒聊] 對於XG挺失望的</a>
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">iamfenixsc</div>
			</div>
		</div>
        <div class="r-list-sep"></div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f1">爆</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1455811061.A.FED.html">[公告] 伺服器狀況詢問/聊天/揪團/抱怨/多功能區</a>
			</div>
			<div class="meta">
				<div class="date"> 2/18</div>
				<div class="author">rainnawind</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"></div>
			<div class="mark">!</div>
			<div class="title">
				<a href="/bbs/LoL/M.1416199565.A.6B5.html">[公告] LoL 英雄聯盟版 板規 (2015/11/01 ver.)</a>
			</div>
			<div class="meta">
				<div class="date">11/17</div>
				<div class="author">NeVerEnouGh</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f3">27</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1447000166.A.74F.html">[電競] 近期賽事</a>
			</div>
			<div class="meta">
				<div class="date">11/09</div>
				<div class="author">fkc</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f3">27</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1447000199.A.74F.html">[電競] 近期賽事</a>
			</div>
			<div class="meta">
				<div class="date">11/10</div>
				<div class="author">abc</div>
			</div>
		</div>
		<div class="r-ent">
			<div class="nrec"><span class="hl f1">爆</span></div>
			<div class="mark">M</div>
			<div class="title">
				<a href="/bbs/LoL/M.1457605826.A.732.html">[電競] 2016 LMS Spring W7D1</a>
			</div>
			<div class="meta">
				<div class="date"> 3/10</div>
				<div class="author">LMSPostBot</div>
			</div>
		</div>
	</div>
</div>
    </body>
</html>
`

var TestContentHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">


<meta name="viewport" content="width=device-width">

<title>Re: [新聞] 5年7起隨機殺人 台灣社會學到甚麼？ - 看板 Gossiping - 批踢踢實業坊</title>
<meta name="robots" content="all">
<meta name="keywords" content="Ptt BBS 批踢踢">
<meta name="description" content="
學不到什麼的啦

要說因為經濟太差

">
<meta property="og:site_name" content="Ptt 批踢踢實業坊">
<meta property="og:title" content="Re: [新聞] 5年7起隨機殺人 台灣社會學到甚麼？">
<meta property="og:description" content="
學不到什麼的啦

要說因為經濟太差

">
<link rel="canonical" href="https://www.ptt.cc/bbs/Gossiping/M.1459688662.A.CFE.html">

<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs-common.css">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs.css" media="screen">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/pushstream.css" media="screen">
<link rel="stylesheet" type="text/css" href="//images.ptt.cc/v2.14/bbs-print.css" media="print">


<script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
<script src="//images.ptt.cc/v2.14/bbs.js"></script>




<script type="text/javascript">

  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-32365737-1']);
  _gaq.push(['_setDomainName', 'ptt.cc']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();

</script>


	</head>
    <body>

<div id="fb-root"></div>
<script>(function(d, s, id) {
var js, fjs = d.getElementsByTagName(s)[0];
if (d.getElementById(id)) return;
js = d.createElement(s); js.id = id;
js.src = "//connect.facebook.net/en_US/all.js#xfbml=1";
fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));</script>

<div id="topbar-container">
	<div id="topbar" class="bbs-content">
		<a id="logo" href="/">批踢踢實業坊</a>
		<span>&rsaquo;</span>
		<a class="board" href="/bbs/Gossiping/index.html"><span class="board-label">看板 </span>Gossiping</a>
		<a class="right small" href="/about.html">關於我們</a>
		<a class="right small" href="/contact.html">聯絡資訊</a>
	</div>
</div>
<div id="navigation-container">
	<div id="navigation" class="bbs-content">
		<a class="board" href="/bbs/Gossiping/index.html">返回看板</a>
		<div class="bar"></div>
		<div class="share">
			<span>分享</span>
			<div class="fb-like" data-send="false" data-layout="button_count" data-width="90" data-show-faces="false" data-href="http://www.ptt.cc/bbs/Gossiping/M.1459688662.A.CFE.html"></div>

			<div class="g-plusone" data-size="medium"></div>
<script type="text/javascript">
window.___gcfg = {lang: 'zh-TW'};
(function() {
var po = document.createElement('script'); po.type = 'text/javascript'; po.async = true;
po.src = 'https://apis.google.com/js/plusone.js';
var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(po, s);
})();
</script>

		</div>
	</div>
</div>
<div id="main-container">
    <div id="main-content" class="bbs-screen bbs-content"><div class="article-metaline"><span class="article-meta-tag">作者</span><span class="article-meta-value">senyek9527 ()</span></div><div class="article-metaline-right"><span class="article-meta-tag">看板</span><span class="article-meta-value">Gossiping</span></div><div class="article-metaline"><span class="article-meta-tag">標題</span><span class="article-meta-value">Re: [新聞] 5年7起隨機殺人 台灣社會學到甚麼？</span></div><div class="article-metaline"><span class="article-meta-tag">時間</span><span class="article-meta-value">Sun Apr  3 21:04:15 2016</span></div>
學不到什麼的啦

要說因為經濟太差

年輕人沒有希望所以製造出反社會人格是社會共業這種事情太複雜

不如在廢死還沒有開始實施前

怪都是廢死的錯比較實際

還沒廢死就這樣

廢死之後一定很誇張

死的如果是你家人你要怎麼想

太冷靜就代表你是廢死

會有那麼多殺人犯都是廢死的錯



大概就是台灣人學到的教訓

<span class="f2">※ 引述《sunbysea (忠言逆耳良藥苦口)》之銘言：
</span><span class="f6">: 影片新聞
</span><span class="f6">: <a href="http://www.ntdtv.com.tw/b5/20160402/video/169005.html" target="_blank" rel="nofollow">http://www.ntdtv.com.tw/b5/20160402/video/169005.html</a>
</span><span class="f6">: 1.媒體來源:
</span><span class="f6">: 新唐人亞太台
</span><span class="f6">: 2.完整新聞標題:
</span><span class="f6">: 5年7起隨機殺人 台灣社會學到甚麼？
</span><span class="f6">: 3.完整新聞內文:
</span><span class="f6">: 【新唐人亞太台 2016 年 04 月 02 日訊】對於隨機殺人，臺大法律系教授李茂
</span><span class="f6">: 生曾說，時間會愈來愈接近、頻率會愈來愈高。而五年內，台灣至少有7起隨機
</span><span class="f6">: 殺人事件，不過這一次，由於小燈泡媽媽的呼籲，喚醒不少台灣人，更理性的面
</span><span class="f6">: 對，從家庭教育開始，從自己做起，用愛以及善良，進一步，改變社會結構。
</span><span class="f6">: 去年5月29日，龔重安隨機殺害8歲女童，震驚社會，事隔不到一年，隨機殺害女
</span><span class="f6">: 童事件，再次發生。
</span><span class="f6">: 近5年，台灣至少有7起隨機殺人事件，其中，孩童就有3位。曾預言「時間會愈
</span><span class="f6">: 來愈接近、頻率會愈來愈高」的台大法律系教授李茂生，進一步指出：「社會結
</span><span class="f6">: 構不變，潛在性的這類人，只會不斷增加。而改變社會結構的第一步，不外就是
</span><span class="f6">: 人心。」
</span><span class="f6">: 小燈泡母親(2016.03.28 )：「我真的很希望，政府各級單位，能夠做些事情，
</span><span class="f6">: 讓媽媽放心帶小孩，或者是讓媽媽放心工作，另外我認為，這樣子的隨機殺人事
</span><span class="f6">: 件，兇嫌基本上在當時是沒有 理智的，這不是靠立甚麼法，怎麼做處治能夠解
</span><span class="f6">: 決這個問題，我還是希望能從根本，從家庭、從教育，來讓這樣子的人，消失在
</span><span class="f6">: 社會上面，我希望我們以後的子子孫孫，都不要再出現這樣子的人。」
</span><span class="f6">: 親眼目睹孩子遭受殺害過程的小燈泡母親，擒著淚，堅忍的說出對社會的期望，
</span><span class="f6">: 倘若，真要杜絕類似案件發生，最根本的，就是家庭以及教育，這些話語，讓每
</span><span class="f6">: 一位父母，有了共鳴。
</span><span class="f6">: 悼念民眾：「這並不是能怪罪任何人，因為，這種無差別的傷害事件，真的很難
</span><span class="f6">: 避免，所以它是社會的根本問題，不要因為這一個，這個事件就否定這個我們所
</span><span class="f6">: 有的人，對於這世界美好的期待，我覺得我們世界上還是有很多好人。」
</span><span class="f6">: 律師 呂秋遠：「在過去來講，可能這兩三年內，已經發生過三件的案件，而每
</span><span class="f6">: 一次，我們發現不管是用死刑，或者用其他的方式做警控跟威脅，似乎對於這樣
</span><span class="f6">: 的人，並沒 有幫助，所以，台灣的民眾慢慢會去省思，到底我們今天要用甚麼
</span><span class="f6">: 樣的方式，來阻止這樣的情況，一而再再而三的發生，而不是單純用理盲，或者
</span><span class="f6">: 激情的方式去解決 這個問題。」
</span><span class="f6">: 案發現場，民眾自發性，送上花束，玩偶，寫上小卡片，一點一點的心意，連結
</span><span class="f6">: 了人性最初的善良與溫暖，此時，收起評論「廢死」或是「反廢死」的聲浪，釋
</span><span class="f6">: 放出的，是更深刻的反省。
</span><span class="f6">: 公益平台文化基金會董事長 嚴長壽(2016.03.29)：「在過去這段時間，台灣的
</span><span class="f6">: 經濟相對的優渥了以後，父母常常，整個社會的環境，都在某種情況，讓孩子有
</span><span class="f6">: 這種機會，不必為自 己負責任，而這些大概都是造成這些現象的重要的原因。
</span><span class="f6">: 從教育的角度，我覺得，一定要教年輕人有生活的能力，我們太過注重在就業的
</span><span class="f6">: 能力，常常忽略掉了，考試 的能力以外，大部分的年輕人都沒有就業能力以外
</span><span class="f6">: 的技術。」
</span><span class="f6">: 有觀光教父之稱的公益平台文化基金會董事長嚴長壽，點出台灣社會正面臨的結
</span><span class="f6">: 構問題，而當教育只是為了升學，為了考試，那麼，擁有考試以外能力的孩子，
</span><span class="f6">: 在這樣的體制下，可能就會被，篩選出去。
</span><span class="f6">: 人本教育基金會執行長 馮喬蘭：「教育要做的事情是，我不管小孩個人條件如
</span><span class="f6">: 何，但我都盡量的帶到你，盡量的讓你覺得被接納，你被肯定，你被當一回事，
</span><span class="f6">: 你被當一個完完整整的個人看 待，假使我們的教育願意這樣子做，老師們可以
</span><span class="f6">: 有這樣子的，你也許沒有辦法一下子甚麼都做到最充分，但有這樣子的轉念，我
</span><span class="f6">: 相信我們可以承接住很多孩子。」
</span><span class="f6">: 小燈泡母親(2016.03.29)：「我真的很希望，她的走，能留下一些什麼，能夠讓
</span><span class="f6">: 一些事情發酵，喚回大家對愛的重視，對社會上許許多多的重視。」
</span><span class="f6">: 的確，對台灣人來說，這起隨機殺人事件，讓人震驚、憤怒與難過，但，小燈泡
</span><span class="f6">: 母親，每一次面對媒體，都不斷的喚醒社會大眾，對家庭、對愛以及關懷的重視
</span><span class="f6">: 。
</span><span class="f6">: 家長 林先生：「我覺得還是從父母對小孩基本的這種，道德教育方面，從小時
</span><span class="f6">: 候培養。」
</span><span class="f6">: 家長 張女士：「以後台灣社會會變成怎樣我們不太清楚，那唯有就是把我們的
</span><span class="f6">: 小孩子教育好，就像那個小燈泡媽媽說的，不要，希望社會上不要再出現這種，
</span><span class="f6">: 令人害怕的 人出現，我們只能管教我們的小孩子，走向比較正向的思維，小孩
</span><span class="f6">: 子多去關心別人，有同理心一點，或許，他對社會的關愛，就會比較多一點。」
</span><span class="f6">: 4月1日凌晨，再過幾個小時，這些追悼物品，環保局就會清空，儘管如此，仍有
</span><span class="f6">: 民眾前來悼念，彎著腰，閱讀卡片上的內容。台灣民眾，正在透過哀悼，學習面
</span><span class="f6">: 對。
</span><span class="f6">: 匯聚民眾愛心的街道，陸續回復樣貌，儘管，這駭人的傷痛，還留在心中，不過
</span><span class="f6">: ，台灣民眾也開始，如同小燈泡的母親一般，認真勇敢面對、去思考，尋求善的
</span><span class="f6">: 力量，從自己做起。一起把社會的洞，理性認真的，補起來。
</span><span class="f6">: 採訪撰稿：李晶晶
</span><span class="f6">: 攝影後製：陳輝模
</span><span class="f6">: 4.完整新聞連結 (或短網址):
</span><span class="f6">: <a href="http://www.ntdtv.com.tw/b5/20160402/video/169005.html" target="_blank" rel="nofollow">http://www.ntdtv.com.tw/b5/20160402/video/169005.html</a>
</span><span class="f6">: 5.備註:
</span>

-----
Sent from JPTT on my Xiaomi Mi 4i.

--
機會就像老二 握緊就會變大
時間就像乳溝 擠擠還是有的

--
<span class="f2">※ 發信站: 批踢踢實業坊(ptt.cc), 來自: 126.148.111.173
</span><span class="f2">※ 文章網址: <a href="https://www.ptt.cc/bbs/Gossiping/M.1459688662.A.CFE.html" target="_blank" rel="nofollow">https://www.ptt.cc/bbs/Gossiping/M.1459688662.A.CFE.html</a>
</span><div class="push"><span class="f1 hl push-tag">噓 </span><span class="f3 hl push-userid">skyexers</span><span class="f3 push-content">: 還沒崩潰完？</span><span class="push-ipdatetime"> 04/03 21:07
</span></div><div class="push"><span class="f1 hl push-tag">→ </span><span class="f3 hl push-userid">hkcdc</span><span class="f3 push-content">: 不判死的確是問題</span><span class="push-ipdatetime"> 04/03 21:15
</span></div><div class="push"><span class="hl push-tag">推 </span><span class="f3 hl push-userid">pictograma</span><span class="f3 push-content">: 一樓崩潰標準示範</span><span class="push-ipdatetime"> 04/03 21:15
</span></div><div class="push"><span class="f1 hl push-tag">→ </span><span class="f3 hl push-userid">drigo</span><span class="f3 push-content">: 學到原來台灣人書唸不夠多, 所以獵廢死</span><span class="push-ipdatetime"> 04/03 21:16
</span></div><div class="push"><span class="f1 hl push-tag">噓 </span><span class="f3 hl push-userid">ainor</span><span class="f3 push-content">: 一堆22K都沒殺人了</span><span class="push-ipdatetime"> 04/03 21:19
</span></div><div class="push"><span class="f1 hl push-tag">→ </span><span class="f3 hl push-userid">dave01</span><span class="f3 push-content">: 沒學到  過一陣子 霉不報導 一堆人就忘了</span><span class="push-ipdatetime"> 04/03 21:30
</span></div><div class="push"><span class="hl push-tag">推 </span><span class="f3 hl push-userid">puorg</span><span class="f3 push-content">: 你太中肯了</span><span class="push-ipdatetime"> 04/03 22:27
</span></div><div class="push"><span class="f1 hl push-tag">→ </span><span class="f3 hl push-userid">xgodtw</span><span class="f3 push-content">: 有點中肯</span><span class="push-ipdatetime"> 04/03 23:51
</span></div></div>

    <div id="article-polling" data-pollurl="/poll/Gossiping/M.1459688662.A.CFE.html?cacheKey=2052-1079006705&offset=5548&offset-sig=352feaabe4bcbcc99dd20155e6321783668b358b" data-longpollurl="/v1/longpoll?id=119b02befb2f87528ab570bf486d9279ccdfbf0d" data-offset="5548"></div>



<div class="bbs-screen bbs-footer-message">本網站已依台灣網站內容分級規定處理。此區域為限制級，未滿十八歲者不得瀏覽。</div>

</div>

    </body>
</html>
`
