package tests

const dummyHTML1 = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<title>W3C QA - Recommended list of Doctype declarations you can use in your Web document</title>
	<meta name="Keywords" content="qa, quality assurance, conformance, validity, test suite, QA, quality, validation, html, validator, tools" />
	<meta name="Description" content="W3C QA - List of valid Doctype declarations you can use in your document" />

	<link rel="schema.DC" href="http://purl.org/dc" />
	<meta name="DC.Subject" lang="en" content="QA, quality, validation, html, validator, tools" />
	<meta name="DC.Title" lang="en" content="How to achieve web standards and quality on your website?" />
	<meta name="DC.Description.Abstract" lang="en" content="A series of things to improve and achieve quality on your website." />
	<meta name="DC.Date.Created" content="2002-04-08" />
	<meta name="DC.Language" scheme="RFC1766" content ="en" />
	<meta name="DC.Creator" content="Karl Dubost" />
	<meta name="DC.Publisher" content="W3C - World Wide Web Consortium - http://www.w3.org" />
	<meta name="DC.Rights" content="http://www.w3.org/Consortium/Legal/copyright-documents-19990405" />
<style type="text/css" media="all">
@import "/QA/2006/01/blogstyle.css";
</style>
</head>
<body>

<!-- Header -->
<div id="banner">
  <h1 id="title">
	<a href="http://www.w3.org/"><img height="48" alt="W3C" id="logo" src="http://www.w3.org/Icons/WWW/w3c_home_nb" /></a><a href="http://www.w3.org/QA/"><img src="http://www.w3.org/QA/2002/12/qa.png" alt="Quality Assurance" /></a> Recommended list of Doctype declarations
</h1>
</div>
<ul class="navbar" id="menu">
		<li><strong><a href="/QA/" title="Quality Assurance Web Site Home">[QA Home]</a></strong></li>

	<li><a href="/QA/Library/" title="Documents and Publications on Web and Quality">Documents</a></li>
	<li><a href="/QA/Tools/" accesskey="3" title="Validators and other Tools">Tools</a></li>
	<li><a href="/QA/IG/#contact">Feedback</a></li>
</ul>
<div id="searchbox">
	<form method="get" action="http://www.google.com/custom" enctype="application/x-www-form-urlencoded">
	<p id="formbox"><input type="text" size="15" class="textfield" name="q" accesskey="E" maxlength="255" /> <input type="submit" class="submitfield" value="Search" id="goButton" name="sa" accesskey="G" /> <input type="hidden" name="cof" value="T:black;LW:72;ALC:#ff3300;L:http://www.w3.org/Icons/w3c_home;LC:#000099;LH:48;BGC:white;AH:left;VLC:#660066;GL:0;AWFID:0b9847e42caf283e;" /><input type="hidden" id="searchW3C" name="sitesearch" checked="checked" value="www.w3.org/QA" /><input type="hidden" name="domains" value="www.w3.org/QA" /></p>
	</form>
	</div>


		<div id="main"><!-- This DIV encapsulates everything in this page - necessary for the positioning -->


<div id="jumpbar"> 
	<h2>Warning</h2> 
	<p>The list is <strong>informative</strong> and
does not try to be exhaustive (there are many other proper declarations
you could use), but it has most of the declarations commonly used on the
Web at the moment.</p>
 </div>

<h2>Recommended Doctype Declarations to  use in your Web document.</h2>

<p>When authoring document is HTML or XHTML, it is important to <a href="http://www.w3.org/QA/Tips/Doctype">Add a Doctype declaration</a>. This makes sure the document will be parsed the same way by different browsers.</p>

<p>The simplest and most reliable doctype declaration to use is the one defined in <a href="/TR/html5">HTML5</a>:</p>

<div class="example">
<pre class="html">
&lt;!DOCTYPE html&gt;
</pre></div>

		  <p>If you need a doctype matching a specific version of (X)HTML, the doctype declaration must be exact (both in spelling and in case) to have the desired effect, which makes it sometimes difficult. To ease the 
work, below is a list of recommended doctype declarations that you can use in your Web documents.</p>

<h2 id="Template">Template</h2>
<p>Use the following markup as a template to create a new HTML document using a proper Doctype declaration. See the <a href="#DTD">list</a> below if you wish to use another DTD.</p>

<div class="example">
<pre class="html">
&lt;!DOCTYPE html&gt;
&lt;html&gt;

&lt;head&gt;
	&lt;title&gt;An HTML standard template&lt;/title&gt;
	&lt;meta charset="utf-8"  /&gt;
&lt;/head&gt;

&lt;body&gt;

 &lt;p&gt;… Your HTML content here …&lt;/p&gt;

&lt;/body&gt;
&lt;/html&gt;
</pre>
</div>




<h2 id="DTD">(X)HTML Doctype Declarations List</h2>

<dl>
<dt class="specification"><a href="http://www.w3.org/TR/html5/">HTML5 and beyond</a></dt>
<dd class="dtd"><pre class="html">&lt;!DOCTYPE HTML&gt;</pre></dd>


<dt class="specification"><a href="http://www.w3.org/TR/html401/">HTML&nbsp;4.01</a></dt>
<dd><dl><dt><a href="http://www.w3.org/TR/html401/strict.dtd">Strict</a></dt>
	<dd class="dtd"><pre class="html">&lt;!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
"http://www.w3.org/TR/html4/strict.dtd"&gt;</pre></dd>
<dt><a href="http://www.w3.org/TR/html401/loose.dtd">Transitional</a></dt>
	<dd class="dtd"><pre class="html">&lt;!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
"http://www.w3.org/TR/html4/loose.dtd"&gt;</pre></dd>

<dt><a href="http://www.w3.org/TR/html401/frameset.dtd">Frameset</a></dt>
	<dd class="dtd"><pre class="html">&lt;!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"
"http://www.w3.org/TR/html4/frameset.dtd"&gt;
	</pre>
</dd>
</dl></dd>
	<dt class="specification"><a href="http://www.w3.org/TR/xhtml1/">XHTML&nbsp;1.0</a></dt>
<dd><dl><dt><a href="http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">Strict</a> (<a href="http://www.w3.org/2010/04/xhtml10-strict.html">quick reference</a>)</dt>
<dd class="dtd"><pre class="html">&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"&gt;</pre></dd>
<dt><a href="http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">Transitional</a></dt>
<dd class="dtd"><pre class="html">&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"&gt;</pre></dd>
<dt><a href="http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">Frameset</a></dt>
<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd"&gt;</pre>
</dd>
</dl></dd>
	<dt class="specification"><a href="http://www.w3.org/TR/xhtml11/">XHTML&nbsp;1.1</a> - <a href="http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" 
"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd"&gt;
	</pre>
</dd>
	<dt class="specification"><a href="http://www.w3.org/TR/2006/WD-xhtml-basic-20060705">XHTML&nbsp;Basic&nbsp;1.1</a> (<a href="http://www.w3.org/2007/07/xhtml-basic-ref.html">quick reference</a>): 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"
"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd"&gt;
	</pre>
</dd>

</dl>

<h2>MathML Doctype Declarations</h2>
<dl>

	<dt class="specification"><a href="http://www.w3.org/TR/mathml2">MathML 2.0</a> - <a href="http://www.w3.org/Math/DTD/mathml2/mathml2.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE math PUBLIC "-//W3C//DTD MathML 2.0//EN"
	"http://www.w3.org/Math/DTD/mathml2/mathml2.dtd"&gt;
	</pre>
</dd>

	<dt class="specification"><a href="http://www.w3.org/pub/WWW/TR/REC-MathML/">MathML 1.01</a> - <a href="http://www.w3.org/Math/DTD/mathml1/mathml.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE math SYSTEM 
	"http://www.w3.org/Math/DTD/mathml1/mathml.dtd"&gt;
	</pre>
</dd>

</dl>

<h2>Compound documents doctype declarations</h2>
<dl>

	<dt class="specification"><a href="http://www.w3.org/TR/XHTMLplusMathMLplusSVG/">XHTML + MathML + SVG</a> - <a href="http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC
"-//W3C//DTD XHTML 1.1 plus MathML 2.0 plus SVG 1.1//EN"
"http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd"&gt;
	</pre>
</dd>

	<dt class="specification">XHTML + MathML + SVG Profile (XHTML as the host language) - <a href="http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC
"-//W3C//DTD XHTML 1.1 plus MathML 2.0 plus SVG 1.1//EN"
"http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd"&gt;
	</pre>
</dd>
	<dt class="specification">XHTML + MathML + SVG Profile (Using SVG as the host) - <a href="http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE svg:svg PUBLIC
"-//W3C//DTD XHTML 1.1 plus MathML 2.0 plus SVG 1.1//EN"
"http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd"&gt;
	</pre>
</dd>
</dl>

<h2>Optional doctype declarations</h2>

<p>Beyond the specificities of (X)HTML processing, Doctype declarations in XML languages are only useful to declare named entities and to facilitate the validation of documents based on DTDs. This means that in many XML languages, doctype declarations are not necessarily useful.</p>

<p>The list below is provided only if you actually need to declare a doctype for these types of documents.</p>

<dl>
	<dt class="specification"><a href="http://www.w3.org/TR/2003/REC-SVG11-20030114/">SVG 1.1 Full</a> - <a href="http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
	"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"&gt;
	</pre>
</dd>
	<dt class="specification"><a href="http://www.w3.org/TR/2001/REC-SVG-20010904/">SVG 1.0</a> - <a href="http://www.w3.org/TR/2001/REC-SVG-20010904/DTD/svg10.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.0//EN"
	"http://www.w3.org/TR/2001/REC-SVG-20010904/DTD/svg10.dtd"&gt;
	</pre>
</dd>
	<dt class="specification"><a href="http://www.w3.org/TR/2003/REC-SVGMobile-20030114/">SVG 1.1 Basic</a> - <a href="http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-basic.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Basic//EN"
	"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-basic.dtd"&gt;
	</pre>
</dd>
	<dt class="specification"><a href="http://www.w3.org/TR/2003/REC-SVGMobile-20030114/">SVG 1.1 Tiny</a> - <a href="http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Tiny//EN"
	"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd"&gt;
	</pre>
</dd>

</dl>

<h2>Historical doctype declarations</h2>

<p>The doctype declarations below are valid, but have mostly an historical value — a doctype declaration of a more recent equivalent ought to be used in their stead.</p>

<dl>
<dt class="specification"><a href="http://www.ietf.org/rfc/rfc1866.txt">HTML&nbsp;2.0</a> - <a href="http://www.w3.org/MarkUp/html-spec/html.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">&lt;!DOCTYPE html PUBLIC "-//IETF//DTD HTML 2.0//EN"&gt;</pre>
</dd>

	<dt class="specification"><a href="http://www.w3.org/TR/REC-html32">HTML&nbsp;3.2</a> - <a href="http://www.w3.org/MarkUp/Wilbur/HTML32.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN"&gt;
</pre>
</dd>
	<dt class="specification"><a href="http://www.w3.org/TR/2000/REC-xhtml-basic-20001219/">XHTML&nbsp;Basic&nbsp;1.0</a> - <a href="http://www.w3.org/TR/xhtml-basic/xhtml-basic10.dtd">DTD</a>: 

</dt>
	<dd class="dtd"><pre class="html">
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.0//EN"
"http://www.w3.org/TR/xhtml-basic/xhtml-basic10.dtd"&gt;
	</pre>
</dd>

</dl>

</div>
<address>
  <a href="http://validator.w3.org/check?uri=referer"><img
	src="http://www.w3.org/Icons/valid-xhtml10" height="31" width="88"
	alt="Valid XHTML 1.0!" /></a>
  <a href="/QA/IG/#contact">Send us Feedback</a>      <br />
  Last updated: $Date: 2016/07/01 13:20:56 $ by $Author: dom $
</address>



<p class="copyright"> <a rel="Copyright"
href="http://www.w3.org/Consortium/Legal/ipr-notice#Copyright">Copyright</a>
&copy; 1994-2006 <a href="http://www.w3.org/"><acronym title="World Wide
Web Consortium">W3C</acronym></a>&reg; (<a
href="http://www.csail.mit.edu/"><acronym title="Massachusetts Institute of
Technology">MIT</acronym></a>, <a href="http://www.ercim.org/"><acronym
title="European Research Consortium for Informatics and
Mathematics">ERCIM</acronym></a>, <a
href="http://www.keio.ac.jp/">Keio</a>), All Rights Reserved. W3C <a
href="http://www.w3.org/Consortium/Legal/ipr-notice#Legal_Disclaimer">liability</a>,
<a
href="http://www.w3.org/Consortium/Legal/ipr-notice#W3C_Trademarks">trademark</a>,
<a rel="Copyright"
href="http://www.w3.org/Consortium/Legal/copyright-documents">document
use</a> and <a rel="Copyright"
href="http://www.w3.org/Consortium/Legal/copyright-software">software
licensing</a> rules apply. Your interactions with this site are in
accordance with our <a
href="http://www.w3.org/Consortium/Legal/privacy-statement#Public">public</a>
and <a
href="http://www.w3.org/Consortium/Legal/privacy-statement#Members">Member</a>
privacy statements. </p>

</body>
</html>"`

const dummyHTML2 = `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<!-- $Id: common.m4,v 1.2 2004/10/02 00:07:56 smoot Exp $ -->
<html xmlns="https://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
<head>
        <meta name="author" content="Smoot Carl-Mitchell"/>

        <title>Home</title>
        <link rel="stylesheet" href="tic.css"/>
</head>

<body>
<!-- Site common elements -->

<!-- Site logo -->
<div class="page_title">
        <img class="tic_logo" src="bigticlogo.png" alt="[TIC Logo]"/>
        <h1 class="title">&nbsp; Home </h1>

</div>

<hr/>

<!-- Site navigation menu -->
<div class="navigation">
        <a href="index.html">Home</a>&nbsp;&bull;
        <a href="bios/index.html">Bios</a>&nbsp;&bull;
        <a href="books/index.html">Books</a>&nbsp;&bull;
        <a href="mailto:tic@tic.com">Contact</a>&nbsp;&bull;
        <a href="opensource/index.html">Open Source Projects</a>&nbsp;&bull;
        <a href="partners/index.html">Partners</a>&nbsp;&bull;
        <a href="rfcs/index.html">RFCs</a>&nbsp;&bull;
        <a href="whitepapers/index.html">Whitepapers</a>
</div>

<hr/>
<p>
TIC is a consulting firm specializing in:
</p>
<ul>
<li>GNU/Linux System Architecture</li>
<li>Unix System Architecture</li>
<li>TCP/IP Network Design</li>
<li>Network Security</li>
<li>Open Source Tool Development</li>
</ul>
<!-- $Id: end.m4,v 1.4 2004/10/02 00:07:57 smoot Exp $ -->

<!-- SiteSearch HtDig -->
<!-- $Id: htdig.m4,v 1.2 2004/10/02 00:07:58 smoot Exp $ -->
<div class="htdig">
        <hr/>
        <form method="post" action="/cgi-bin/htsearch">
                <p>
                Search TIC:
                <input type="text" size="30" name="words" value=""/>
                <input type="submit" value="Search"/>
                </p>

                <p>
                Match: <select name="method">
                        <option value="and">All</option>
                        <option value="or">Any</option>
                        <option value="boolean">Boolean</option>
                </select>

                Format: <select name="format">
                        <option value="builtin-long">Long</option>
                        <option value="builtin-short">Short</option>
                </select>

                Sort by: <select name="sort">
                        <option value="score">Score</option>
                        <option value="time">Time</option>
                        <option value="title">Title</option>
                        <option value="revscore">Reverse Score</option>
                        <option value="revtime">Reverse Time</option>
                        <option value="revtitle">Reverse Title</option>
                </select>

                <input type="hidden" name="config" value="htdig"/>
                <input type="hidden" name="restrict" value=""/>
                <input type="hidden" name="exclude" value=""/>
                </p>

        </form>
        <hr/>
</div>


<div>
        <a href="https://www.apache.org">
                <img src="apache_pb.png" alt="[Apache Logo]"/>
        </a>

        <a href="https://www.htdig.org">
                <img src="/htdig/htdig.gif" alt="ht://Dig"/>
        </a>
        <a href="https://validator.w3.org/check?uri=referer">
                <img src="https://www.w3.org/Icons/valid-xhtml10"
                alt="Valid XHTML 1.0!" height="31" width="88" />
        </a>
</div>
<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-489884-12");
pageTracker._initData();
pageTracker._trackPageview();
</script>
</body>
</html>`

const dummyHTML3 = `<!DOCTYPE html>
<!-- $Id: common.m4,v 1.2 2004/10/02 00:07:56 smoot Exp $ -->
<html xmlns="https://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
<head>
        <meta name="author" content="Smoot Carl-Mitchell"/>

        <title>Old site in the web</title>
        <link rel="stylesheet" href="tic.css"/>
</head>

<body>
<!-- Site common elements -->

<!-- Site logo -->
<div class="page_title">
        <img class="tic_logo" src="bigticlogo.png" alt="[TIC Logo]"/>
        <h1 class="title">&nbsp; Home </h1>

</div>

<hr/>

<!-- Site navigation menu -->
<div class="navigation">
        <a href="index.html">Home</a>&nbsp;&bull;
        <a href="bios/index.html">Bios</a>&nbsp;&bull;
        <a href="books/index.html">Books</a>&nbsp;&bull;
        <a href="mailto:tic@tic.com">Contact</a>&nbsp;&bull;
        <a href="opensource/index.html">Open Source Projects</a>&nbsp;&bull;
        <a href="partners/index.html">Partners</a>&nbsp;&bull;
        <a href="rfcs/index.html">RFCs</a>&nbsp;&bull;
        <a href="whitepapers/index.html">Whitepapers</a>
</div>

<hr/>
<p>
TIC is a consulting firm specializing in:
</p>
<ul>
<li>GNU/Linux System Architecture</li>
<li>Unix System Architecture</li>
<li>TCP/IP Network Design</li>
<li>Network Security</li>
<li>Open Source Tool Development</li>
</ul>
<!-- $Id: end.m4,v 1.4 2004/10/02 00:07:57 smoot Exp $ -->

<!-- SiteSearch HtDig -->
<!-- $Id: htdig.m4,v 1.2 2004/10/02 00:07:58 smoot Exp $ -->
<div class="htdig">
        <hr/>
        <form method="post" action="/cgi-bin/htsearch">
                <p>
                Search TIC:
                <input type="text" size="30" name="words" value=""/>
                <input type="submit" value="Search"/>
                </p>

                <p>
                Match: <select name="method">
                        <option value="and">All</option>
                        <option value="or">Any</option>
                        <option value="boolean">Boolean</option>
                </select>

                Format: <select name="format">
                        <option value="builtin-long">Long</option>
                        <option value="builtin-short">Short</option>
                </select>

                Sort by: <select name="sort">
                        <option value="score">Score</option>
                        <option value="time">Time</option>
                        <option value="title">Title</option>
                        <option value="revscore">Reverse Score</option>
                        <option value="revtime">Reverse Time</option>
                        <option value="revtitle">Reverse Title</option>
                </select>

                <input type="hidden" name="config" value="htdig"/>
                <input type="hidden" name="restrict" value=""/>
                <input type="hidden" name="exclude" value=""/>
                </p>

        </form>
        <hr/>
</div>


<div>
        <a href="https://www.apache.org">
                <img src="apache_pb.png" alt="[Apache Logo]"/>
        </a>

        <a href="https://www.htdig.org">
                <img src="/htdig/htdig.gif" alt="ht://Dig"/>
        </a>
        <a href="https://validator.w3.org/check?uri=referer">
                <img src="https://www.w3.org/Icons/valid-xhtml10"
                alt="Valid XHTML 1.0!" height="31" width="88" />
        </a>
</div>
<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-489884-12");
pageTracker._initData();
pageTracker._trackPageview();
</script>
</body>
</html>`

const dummyHTML4 = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"  "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" class="">

<head>
    <meta charset="utf-8">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="robots" content="noindex, nofollow">

            <meta name="viewport" content="width=device-width,initial-scale=1"/>
    <title>Anmelden</title>
    <link rel="icon" type="image/x-icon" href="data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBlXWCgZs1goGeBYKBn0WCgZ91goGelYKBnFWCgZf1goGR5YKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBlGWCgZuVgoGehYKBnmWCgZvVgoGWRYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkaWCgZw1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGYJYKBkAWCgZAFgoGQBYKBkAWCgZW1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGcRYKBkUAAAAAAAAAABYKBkAWCgZAFgoGQBYKBkAWCgZHVgoGe1YKBn/WCgZ/1goGf9YKBmcWCgZQlgoGSNYKBkrWCgZVVgoGaRYKBn/WCgZ/1goGbZYKBkAWCgZAFgoGRlYKBn/WCgZ/1goGf9YKBn/WCgZ6VgoGaRYKBnIWCgZ/1goGf8AAAAAAAAAAFgoGQBYKBkAWCgZAFgoGQVYKBnjWCgZ/1goGf9YKBn/WCgZb1goGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGStYKBnIWCgZ/1goGctYKBkAWCgZfFgoGf9YKBn/WCgZ/1goGdpYKBkNWCgZAFgoGQBYKBlcWCgZxAAAAAAAAAAAWCgZAFgoGQBYKBkAWCgZk1goGf9YKBn/WCgZ/1goGcpYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBmxWCgZ/1goGcVYKBngWCgZ/1goGf9YKBn/WCgZOVgoGQBYKBkAWCgZAFgoGQBYKBkAAAAAAAAAAABYKBkAWCgZAFgoGRFYKBn/WCgZ/1goGf9YKBn/WCgZVlgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBmyWCgZ/1goGf9YKBn/WCgZ/1goGdpYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQAAAAAAAAAAAFgoGQBYKBkAWCgZUlgoGf9YKBn/WCgZ/1goGfVYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBngWCgZ/1goGf9YKBn/WCgZi1goGQBYKBkAWCgZAFgoGQBYKBkAWCgZAAAAAAAAAAAAWCgZclgoGXJYKBnJWCgZ/1goGf9YKBn/WCgZ7VgoGW1YKBlzWCgZblgoGRNYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZBVgoGe9YKBn/WCgZ/1goGf9YKBlMWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAAAAAAAAAAABYKBnjWCgZ/1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ6FgoGQFYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBk8WCgZ/1goGf9YKBn/WCgZ+1goGQhYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQAAAAAAAAAAAFgoGRdYKBmEWCgZ6VgoGf9YKBn/WCgZ/1goGdZYKBmJWCgZilgoGYpYKBmNWCgZA1goGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGW1YKBn/WCgZ/1goGf9YKBn8WCgZPlgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQBYKBnGWCgZ/1goGf9YKBn/WCgZlFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZpFgoGf9YKBn/WCgZ/1goGf9YKBneWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAAAAAAAAAAABYKBlNWCgZ6VgoGfpYKBn/WCgZ/1goGf9YKBnyWCgZ21goGdxYKBncWCgZ2lgoGWRYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBnXWCgZ/1goGf9YKBn/WCgZ/1goGf9YKBlZWCgZAFgoGQBYKBkAWCgZAFgoGQAAAAAAAAAAAFgoGRxYKBn/WCgZ/1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGQZYKBkAWCgZAFgoGQBYKBkAWCgZFFgoGf5YKBn/WCgZ/1goGfFYKBmQWCgZ/1goGelYKBkAWCgZAFgoGQBYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQ5YKBm0WCgZ/1goGf9YKBn/WCgZ51goGS1YKBkwWCgZMFgoGTFYKBk4WCgZBVgoGQBYKBkAWCgZAFgoGQBYKBlLWCgZ/1goGf9YKBn/WCgZ2VgoGQBYKBnTWCgZ/1goGVVYKBkAWCgZAFgoGQBYKBkAAAAAAAAAAABYKBkAWCgZAFgoGWZYKBn/WCgZ/1goGf9YKBn6WCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGYJYKBn/WCgZ/1goGf9YKBmZWCgZAFgoGX1YKBn/WCgZ0lgoGQBYKBkAWCgZAFgoGQAAAAAAAAAAAFgoGQBYKBkAWCgZG1goGf9YKBn/WCgZ/1goGf9YKBlZWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZvlgoGf9YKBn/WCgZ/1goGS9YKBkAWCgZPlgoGf9YKBn/WCgZMFgoGQBYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQBYKBkAWCgZq1goGf9YKBn/WCgZ/1goGbpYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQlYKBn5WCgZ/1goGf9YKBnWWCgZAFgoGQBYKBkQWCgZ+VgoGf9YKBmEWCgZAFgoGQBYKBkAAAAAAAAAAABYKBkAWCgZAFgoGQBYKBkYWCgZ/1goGf9YKBn/WCgZ/1goGR1YKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZUFgoGf9YKBn/WCgZ/1goGVJYKBkAWCgZAFgoGQBYKBnfWCgZ/1goGcZYKBkAWCgZAFgoGQAAAAAAAAAAAFgoGQBYKBkAWCgZAFgoGQBYKBlgWCgZ/1goGf9YKBn/WCgZvVgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBm/WCgZ/1goGf9YKBnPWCgZAFgoGQBYKBkAWCgZAFgoGdlYKBn/WCgZ+FgoGQxYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBmHWCgZ/1goGf9YKBn/WCgZhFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZQlgoGf9YKBn/WCgZ91goGRhYKBkAWCgZAFgoGQBYKBkAWCgZ4FgoGf9YKBn/WCgZMlgoGQBYKBkAAAAAAAAAAABYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBl1WCgZ/1goGf9YKBn/WCgZsFgoGUBYKBkXWCgZI1goGWxYKBn/WCgZ/1goGeJYKBkqWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBn5WCgZ/1goGf9YKBlEWCgZAFgoGQAAAAAAAAAAAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBk1WCgZsFgoGf9YKBn/WCgZ/1goGf9YKBn/WCgZ/1goGflYKBmKWCgZCVgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGXJYKBmxWCgZqVgoGRdYKBkAWCgZAAAAAAAAAAAAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZNlgoGXBYKBmbWCgZrVgoGZpYKBljWCgZHlgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAWCgZAFgoGQBYKBkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA////////////////////////////AH4H/AA8AfgAGAHwPwgZ8H/AP+B/4H/g//B/gA/gf4AH4H+AB+B/4P/gf4AH4D+AA8A/wAPBH+D/wR/gf8EP8H+DD/A/g4/4P4eH/B8Hh/4AD4f/AB+H/8B///////////////////////8=" />
    <script>
      function appendStyle(style) {
        const head = document.getElementsByTagName('HEAD')[0];
        const link = document.createElement('link');
        link.rel = 'stylesheet';
        link.type = 'text/css';
        link.href = "/auth/resources/mjlzb/login/advanzia-de/" + style;
        head.appendChild(link);
      }
      function mapBrandStyle(client) {
        const path = 'lib/backbase-identity/styles/';
        if (client === 'advanzia-web-client') {
          return path + 'identity-B2C-brand.css';
        }
        return path + 'identity-B2B-brand.css';
      }
    </script>
            <link href="/auth/resources/mjlzb/login/advanzia-de/lib/zocial/zocial.css" rel="stylesheet" />
            <link href="/auth/resources/mjlzb/login/advanzia-de/lib/backbase-theme/dist/backbase-theme.css" rel="stylesheet" />
        <link href="/auth/resources/mjlzb/login/advanzia-de/lib/backbase-identity/styles/identity-B2C-brand.css" rel="stylesheet" />
            <script src="/auth/resources/mjlzb/login/advanzia-de/lib/backbase-identity/scripts/load-background.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/dropdownHandler.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/modal.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/selectDeviceForm.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/checkDeviceForm.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/login.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/otpFormHandler.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/notificationHandler.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/resetPwdHandler.js" type="text/javascript"></script>
            <script src="/auth/resources/mjlzb/login/advanzia-de/scripts/loadingButtonHandler.js" type="text/javascript"></script>
</head>
<body onload="loadBackground();loginInit();">
  <div class="bb-stack bb-stack--wrap vh-100 ad-no-overflow">
    <div class="bb-stack__item bb-stack__item--fill identity-bg-layout col-6 d-md-block d-none"></div>
    <div class="bb-stack__break"></div>
    <div class="bb-stack__item bb-stack__item--fill identity-container col-12 col-md-6 ad-overflow">
      <div class="identity-container__panel identity-container__panel--animated">
        <div id="kc-header" class="d-none">
          <div id="kc-header-wrapper" class="">Advanzia Mobile Customer (German)</div>
        </div>
        <div id="form-container" class="identity-container__form ">
          <div class="bb-stack">
            <div class="bb-stack__item identity-logo bb-block bb-block--lg"></div>
          </div>
          <header class="">
            <h1 id="kc-page-title">            Willkommen bei Ihrem Kundenportal

</h1>
          </header>

        <div class="bb-subtitle bb-block bb-block--lg font-primary font-regular">
            <span>Um sich sicher anzumelden, geben Sie unten Ihre Daten ein</span>
        </div>


          <div id="adv-info" class="alert alert-info ad-alert-info ad-mb-4 ad-hidden">
            <div class="alert-body">
              <i aria-hidden="true" class="bb-icon bb-icon--md bb-icon-info mr-2"></i>
              <div class="alert-content">
                <span class="small"><label class="kc-feedback-text font-primary font-semibold mb-0">Sie sind bereits registriert.</label></span>
              </div>
            </div>
          </div>

    <div id="kc-form">
        <script>
            document.getElementById('form-container').classList.add('ad-justify-start', 'ad-height-auto');
        </script>
        <div id="kc-form-wrapper">
                <div id="ad-login-usr-pwd-no-input-error" class="alert alert-danger" hidden>
                    <div class="alert-body">
                        <i aria-hidden="true" class="bb-icon bb-icon--md bb-icon-error-outline mr-2"></i>
                        <div class="alert-content">
                        <span class="small"><label class="kc-feedback-text font-primary font-semibold mb-0">Füllen Sie bitte alle Felder aus</label></span>
                        </div>
                    </div>
                </div>
                <form id="kc-form-login" action="https://identity.gebuhrenfrei.com/auth/realms/advanzia-deu/login-actions/authenticate?session_code=LlwrVaG2eUIsOnQlAftecRX_syoxDxc0p7Ij13HPa9k&amp;execution=6faf0a66-c53d-4fbe-83f2-5abcb2830039&amp;client_id=advanzia-web-client&amp;tab_id=gxk9EejpFU0" method="post">
                    <div class="form-group bb-form-field bb-form-field--md">
                        <label for="username" class="label font-primary font-semibold">Benutzername</label>
                        <input
                            id="username" class=" form-control" name="username" value=""  type="text" autocomplete="off"
                            onfocus="loginController.clearValidationError(this)"
                        />
                    </div>

                    <div class="form-group bb-form-field bb-form-field--md ad-mb-4">
                        <label for="password" class="label font-primary font-semibold">
                            Passwort
                        </label>
                        <div class="input-group">
                            <input id="password" class=" form-control" name="password" type="password" autocomplete="off"
                                onfocus="loginController.clearValidationError(this)"
                            />
                            <button id="ad-password-toggle-button" type="button" class="ad-password-toggle-button "
                                    onclick="loginController.togglePwdVisibility();"
                            >
                                <i id="ad-password-visible-icon-off" class="bb-icon bb-icon-visibility-off bb-icon--md ad-hidden" role="img" aria-hidden="true"></i>
                                <i id="ad-password-visible-icon-on" class="bb-icon bb-icon-visibility bb-icon--md" role="img" aria-hidden="true"></i>
                            </button>
                        </div>
                    </div>

                    <div class="form-group bb-form-field bb-form-field--md bb-block bb-block--xl">
                        <div id="kc-form-buttons" class="bb-button-bar bb-button-bar--reverse">
                            <button
                                class="adv-button ad-button-basic font-primary font-semibold bb-button-bar__button btn btn-primary  "
                                name="login" id="kc-login" type="submit" value="Anmelden"
                                onclick="return loginController.validate()">
                                Anmelden
                                <svg id="spinner" aria-label="Loading indicator" class="bb-loading-indicator__circle ad-loading-indicator__circle ad-hidden" preserveAspectRatio="xMinYMin meet"><circle class="bb-loading-indicator__path" cx="50%" cy="50%" fill="none" r="40%" stroke="currentColor"></circle></svg>
                            </button>
                        </div>
                    </div>

                    <div>
                        <hr class="bb-block bb-block--xl" />
                        <div class="bb-highlight bb-block bb-block--md font-primary font-semibold-md">Probleme beim Anmelden?</div>
                        <div class=" bb-subtitle bb-block bb-block--md ad-mb-1">
                            <span><a class="adv-link font-primary font-regular" href="/auth/realms/advanzia-deu/login-actions/reset-credentials?client_id=advanzia-web-client&amp;tab_id=gxk9EejpFU0">Passwort vergessen?</a></span>
                        </div>
                        <div class=" bb-subtitle bb-block bb-block--md ad-mb-1">
                            <span><a id="registrationLink" class="adv-link font-primary font-regular" href="">Keine Login-Daten? Jetzt registrieren</a></span>
                        </div>
                        <div class=" bb-subtitle bb-block bb-block--md">
                            <span><a id="contactUsLink" class="adv-link font-primary font-regular" href="">Kontaktieren Sie uns</a></span>
                        </div>
                    </div>
                    <hr class="bb-block bb-block--xl" />
                    <div class="ad-mb-small-4"><a id="imprintLink" href="" class="adv-link font-primary font-small" target="_blank" rel=noopener>Impressum</a></div>
                    <div class="ad-mb-small-4"><a id="dataInfoLink" href="" class="adv-link font-primary font-small" target="_blank" rel=noopener>Datenschutz</a></div>
                    <div class="ad-h-30"><a id="legalInfoLink" href="" class="adv-link font-primary font-small" target="_blank" rel=noopener>Rechtliche Informationen</a></div>
                </form>
                <script>
                    loginController.setExternalLinks("https://mein.gebuhrenfrei.com/retail-app-de/redirect");
                </script>
        </div>
    </div>
    <script>
        const currentHref = new URL(location.href);
        const info = currentHref.searchParams.get('adv-info');
        if (info) {
            document.querySelector('#adv-info').classList.toggle('ad-hidden');
        }
    </script>


            <div id="kc-info" class="">
              <div id="kc-info-wrapper" class="">

              </div>
            </div>
        </div>
      </div>
    </div>
  </div>

  <div id="bb-modal" hidden>
    <div aria-hidden="true" class="modal-backdrop fade show"></div>
    <div role="dialog" tabindex="-1" class="modal fade show"  style="display: block" aria-modal="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header" id="bb-modal-header"></div>
                <div class="modal-body" id="bb-modal-body"></div>
            </div>
        </div>
    </div>
  </div>
</body>
</html>`
