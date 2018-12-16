package godoc

import "testing"

func TestParseDoc(t *testing.T) {}

func TestParseSearchResult(t *testing.T) {

}

const goDocHTML = `
<!DOCTYPE html><html lang="en">
<head profile="http://a9.com/-/spec/opensearch/1.1/">
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link href="/-/bootstrap.min.css?v=8bec1bba3e23ecba22cffb197a2d440af410b15d" rel="stylesheet">
  <link href="/-/site.css?v=7d81f4104c89dbe376345f6bfe3e62b4e40d3d06" rel="stylesheet">
  <title>antchfx/xquery - GoDoc</title><meta name="robots" content="NOINDEX">
</head>
<body>
<nav class="navbar navbar-default" role="navigation">
  <div class="container">
  <div class="navbar-header">
    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
      <span class="sr-only">Toggle navigation</span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
    </button>
    <a class="navbar-brand" href="/"><strong>GoDoc</strong></a>
  </div>
  <div class="collapse navbar-collapse">
    <ul class="nav navbar-nav">
        <li><a href="/">Home</a></li>
        <li><a href="/-/about">About</a></li>
    </ul>
    <form class="navbar-nav navbar-form navbar-right" id="x-search" action="/" role="search"><input class="form-control" id="x-search-query" type="text" name="q" placeholder="Search"></form>
  </div>
</div>
</nav>

<div class="container">

  <div class="well">

  <form>
    <div class="input-group">
      <input class="form-control" name="q" autofocus="autofocus" value="antchfx/xquery" placeholder="Search for package by import path or keyword." type="text">
      <span class="input-group-btn">
        <button class="btn btn-default" type="submit">Go!</button>
      </span>
    </div>
  </form>

  </div>
  <p>Try this search on <a href="https://go-search.org/search?q=antchfx%2fxquery">Go-Search</a>
  or <a href="https://github.com/search?q=antchfx%2fxquery+language:go">GitHub</a>.


  <table class="table table-condensed">
    <thead><tr><th>Path</th><th>Synopsis</th></tr></thead>
    <tbody>
      <tr><td>

        <a href="/github.com/antchfx/xquery/html">github.com/antchfx/xquery/html</a>
          <ul class="list-inline">
            <li class="additional-info">3 imports</li>

            <li class="additional-info">· 81 stars</li>
          </ul>

      <td class="synopsis">Package htmlquery provides extract data from HTML documents using XPath expression.</td></tr>

      <tr><td>

        <a href="/github.com/antchfx/xquery/xml">github.com/antchfx/xquery/xml</a>
          <ul class="list-inline">
            <li class="additional-info">2 imports</li>

            <li class="additional-info">· 88 stars</li>
          </ul>

      <td class="synopsis">Package xmlquery provides extract data from XML documents using XPath expression.</td></tr>

      <tr><td>

        <a href="/github.com/Jordanzuo/goutil/xmlUtil">github.com/Jordanzuo/goutil/xmlUtil</a>
          <ul class="list-inline">
            <li class="additional-info">1 imports</li>

            <li class="additional-info">· 4 stars</li>
          </ul>

      <td class="synopsis">xml操作工具类： 此操作工具类来源于:https:/&#8203;/&#8203;github.com/&#8203;antchfx/&#8203;xquery 根据实际情况。我去掉了对golang.org/&#8203;x/&#8203;net/&#8203;html/&#8203;charset的依赖，并添加了各种xml加载函数</td></tr>
    </tbody>
  </table>



</div>
<div id="x-footer" class="clearfix">
  <div class="container">
    <a href="https://github.com/golang/gddo/issues">Website Issues</a>
    <span class="text-muted">|</span> <a href="https://golang.org/">Go Language</a>
    <span class="pull-right"><a href="#">Back to top</a></span>
  </div>
</div>

<div id="x-shortcuts" tabindex="-1" class="modal">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
          <h4 class="modal-title">Keyboard shortcuts</h4>
        </div>
        <div class="modal-body">
          <table>
          <tr><td align="right"><b>?</b></td><td> : This menu</td></tr>
          <tr><td align="right"><b>/</b></td><td> : Search site</td></tr>
          <tr class="text-muted"><td align="right"><b>f</b></td><td> : Jump to identifier</td></tr>
          <tr><td align="right"><b>g</b> then <b>g</b></td><td> : Go to top of page</td></tr>
          <tr><td align="right"><b>g</b> then <b>b</b></td><td> : Go to end of page</td></tr>
          <tr class="text-muted"><td align="right"><b>g</b> then <b>i</b></td><td> : Go to index</td></tr>
          <tr class="text-muted"><td align="right"><b>g</b> then <b>e</b></td><td> : Go to examples</td></tr>
          </table>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>
<script src="/-/jquery-2.0.3.min.js?v=fbf9c77d0c4e3c34a485980c1e5316b6212160c8"></script>
<script src="/-/bootstrap.min.js?v=5ada7c103fc1deabc925cc1fdbbb6e451c21fc70"></script>
<script src="/-/site.js?v=371de731c18d91c499d90b1ab0bf39ecf66d6cf7"></script>
<script type="text/javascript">
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-11222381-8']);
  _gaq.push(['_trackPageview']);
  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
</script>
</body>
</html>
`
