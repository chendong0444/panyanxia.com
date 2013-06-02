package hello

import (
	"html/template"

	"net/http"
)

func init() {
	http.HandleFunc("/proxy", proxyHandler)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	if err := proxyTemplate.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//http://www.google.co.jp
}

var proxyTemplate = template.Must(template.New("response").Parse(proxyTemplateHTML))

const proxyTemplateHTML = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en-us" slick-uniqueid="3">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
   <title>盘眼下 盘点时事 放眼天下</title>
   <meta name="author" content="panyanxia">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <link rel="start" href="./index.html">
   <link rel="stylesheet" href="./css/bootstrap.css" type="text/css">   
   <link rel="stylesheet" href="./css/bootstrap-responsive.css" type="text/css">
</head>
<body>
<!-- NAVBAR
    ================================================== -->
<div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container">
          <button type="button" class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="brand" href="#">盘眼下</a>
          <div class="nav-collapse collapse">
            <ul class="nav">
              <li><a href="/">Home</a></li>
              <li class="active"><a href="#">About</a></li>
            </ul>
          </div><!--/.nav-collapse -->
        </div>
      </div>
    </div>
<div>


<div class="container" style="width:95%">
	<div class="hero-unit">
	Url:
	<input type="text" onfocus="this.select()" onmouseover="this.focus()" value="http://www.panyanxia.com" style="width:70%" maxlength="255" size="80" name="url" id="url"/>
    <input type="button" value="Go" id="btnGo" onclick="$('#iframe1').attr('src','./fetch?url='+$('#url').val())"/>
	<input type="button" value="charset" id="btnCharset" onclick="document.charset='gb2312'"/>
	<br/>
	Img:
	<input type="text" onfocus="this.select()" onmouseover="this.focus()" value="http://www.panyanxia.com/images/shawshank.jpg" style="width:70%" maxlength="255" size="80" name="img" id="img"/>
    <input type="button" value="Show" id="btnShow" onclick="$('#img1').attr('src',$('#img').val())"/>
	</div>
	<iframe id="iframe1" src="" frameborder="0"  width="100%" height="600px" ></iframe>
	<img id="img1" src="" width="285px">
</div>

<div id="footer">
    <div class="container">
        <p class="muted credit">Copyright © 1981-2013 PanYanXia</p>
   </div>
</div>
<!--hidden field-->
<input id="page" name="page" value="0" type="hidden"/>
<!-- Placed at the end of the document so the pages load faster -->
<script src="http://code.jquery.com/jquery-1.9.1.min.js"></script>
<script type="text/javascript" src="./js/bootstrap.js"></script>
<script type="text/javascript" src="./js/main.js"></script>
</body>

</html>

`
