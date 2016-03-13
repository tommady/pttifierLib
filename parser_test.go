package pttifierLib

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const testHTML = `

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
				<a class="btn wide disabled">下頁 &rsaquo;</a>
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
				<div class="author">fkc</div>
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

func TestRuleParsing(t *testing.T) {
	t.Log("going on single rules then go multi-rules combination.")
	root, _ := html.Parse(strings.NewReader(testHTML))

	parseRule := Rule{"失望", ""}
	expectResult := Result{
		"https://www.ptt.cc/bbs/LoL/M.1457622551.A.650.html",
		"[閒聊] 對於XG挺失望的",
		"iamfenixsc",
		"3/10",
	}

	if reciveResults, err := parseRule.Parsing(root, strings.Contains); err == nil {
		if len(reciveResults) == 0 {
			t.Errorf("Expect find 1 result but found 0")
		} else {
			for _, result := range reciveResults {
				if result.Title != expectResult.Title {
					t.Errorf("Expected title:%s, but got:%s", expectResult.Title, result.Title)
				}

				if result.URL != expectResult.URL {
					t.Errorf("Expected URL:%s, but got:%s", expectResult.URL, result.URL)
				}

				if result.Author != expectResult.Author {
					t.Errorf("Expected Author:%s, but got:%s", expectResult.Author, result.Author)
				}

				if result.Date != expectResult.Date {
					t.Errorf("Expected Date:%s, but got:%s", expectResult.Date, result.Date)
				}
			}
		}
	} else {
		t.Errorf("error: %v", err)
	}

}
