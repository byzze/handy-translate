(()=>{var xe=Object.defineProperty;var v=(e,t)=>{for(var n in t)xe(e,n,{get:t[n],enumerable:!0})};var O={};v(O,{SetText:()=>Me,Text:()=>De});var Ce="useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict";var g=(e=21)=>{let t="",n=e;for(;n--;)t+=Ce[Math.random()*64|0];return t};var ve=window.location.origin+"/wails/runtime",l={Call:0,Clipboard:1,Application:2,Events:3,ContextMenu:4,Dialog:5,Window:6,Screens:7,System:8,Browser:9},L=g();function We(e,t,n,o){let i=new URL(ve);i.searchParams.append("object",e),i.searchParams.append("method",t);let r={headers:{}};return n&&(r.headers["x-wails-window-name"]=n),o&&i.searchParams.append("args",JSON.stringify(o)),r.headers["x-wails-client-id"]=L,new Promise((u,p)=>{fetch(i,r).then(s=>{if(s.ok)return s.headers.get("Content-Type")&&s.headers.get("Content-Type").indexOf("application/json")!==-1?s.json():s.text();p(Error(s.statusText))}).then(s=>u(s)).catch(s=>p(s))})}function a(e,t){return function(n,o=null){return We(e,n,t,o)}}var G=a(l.Clipboard),be=0,Se=1;function Me(e){G(be,{text:e})}function De(){return G(Se)}var I={};v(I,{Hide:()=>Ee,Quit:()=>Re,Show:()=>ye});var k=a(l.Application),N={Hide:0,Show:1,Quit:2};function Ee(){k(N.Hide)}function ye(){k(N.Show)}function Re(){k(N.Quit)}var A={};v(A,{GetAll:()=>Ne,GetCurrent:()=>ze,GetPrimary:()=>Ie});var z=a(l.Screens),Le=0,Oe=1,ke=2;function Ne(){return z(Le)}function Ie(){return z(Oe)}function ze(){return z(ke)}var T={};v(T,{IsDarkMode:()=>Pe});var Ae=a(l.System),Te=0;function Pe(){return Ae(Te)}var P={};v(P,{OpenURL:()=>He});var Be=a(l.Browser),je=0;function He(e){Be(je,{url:e})}var Ue=a(l.Call),b=0,f=new Map;function Fe(){let e;do e=g();while(f.has(e));return e}function Z(e,t,n){let o=f.get(e);o&&(n?o.resolve(JSON.parse(t)):o.resolve(t),f.delete(e))}function Y(e,t){let n=f.get(e);n&&(n.reject(t),f.delete(e))}function S(e,t){return new Promise((n,o)=>{let i=Fe();t=t||{},t["call-id"]=i,f.set(i,{resolve:n,reject:o}),Ue(e,t).catch(r=>{o(r),f.delete(i)})})}function Q(e){return S(b,e)}function X(e,...t){if(typeof e!="string"||e.split(".").length!==3)throw new Error("CallByName requires a string in the format 'package.struct.method'");let n=e.split(".");return S(b,{packageName:n[0],structName:n[1],methodName:n[2],args:t})}function V(e,...t){return S(b,{methodID:e,args:t})}function q(e,t,...n){return S(b,{packageName:"wails-plugins",structName:e,methodName:t,args:n})}var Ge=0,Ze=1,Ye=2,Qe=3,Xe=4,Ve=5,qe=6,Je=7,_e=8,Ke=9,$e=10,et=11,tt=12,nt=13,ot=14,it=15,rt=16,lt=17,at=18,st=19,ut=20,ct=21,dt=22,ft=23,mt=24,pt=25,wt=26,gt=27,ht=28,xt=29;function J(e){let t=a(l.Window,e);return{Center:()=>void t(Ge),SetTitle:n=>void t(Ze,{title:n}),Fullscreen:()=>void t(Ye),UnFullscreen:()=>void t(Qe),SetSize:(n,o)=>t(Xe,{width:n,height:o}),Size:()=>t(Ve),SetMaxSize:(n,o)=>void t(qe,{width:n,height:o}),SetMinSize:(n,o)=>void t(Je,{width:n,height:o}),SetAlwaysOnTop:n=>void t(_e,{alwaysOnTop:n}),SetRelativePosition:(n,o)=>t(Ke,{x:n,y:o}),RelativePosition:()=>t($e),Screen:()=>t(et),Hide:()=>void t(tt),Maximise:()=>void t(nt),Show:()=>void t(st),Close:()=>void t(ut),ToggleMaximise:()=>void t(it),UnMaximise:()=>void t(ot),Minimise:()=>void t(rt),UnMinimise:()=>void t(lt),Restore:()=>void t(at),SetBackgroundColour:(n,o,i,r)=>void t(ct,{r:n,g:o,b:i,a:r}),SetResizable:n=>void t(dt,{resizable:n}),Width:()=>t(ft),Height:()=>t(mt),ZoomIn:()=>void t(pt),ZoomOut:()=>void t(wt),ZoomReset:()=>void t(gt),GetZoomLevel:()=>t(ht),SetZoomLevel:n=>void t(xt,{zoomLevel:n})}}var Ct=a(l.Events),vt=0,B=class{constructor(t,n,o){this.eventName=t,this.maxCallbacks=o||-1,this.Callback=i=>(n(i),this.maxCallbacks===-1?!1:(this.maxCallbacks-=1,this.maxCallbacks===0))}},M=class{constructor(t,n=null){this.name=t,this.data=n}},c=new Map;function D(e,t,n){let o=c.get(e)||[],i=new B(e,t,n);return o.push(i),c.set(e,o),()=>Wt(i)}function _(e,t){return D(e,t,-1)}function K(e,t){return D(e,t,1)}function Wt(e){let t=e.eventName,n=c.get(t).filter(o=>o!==e);n.length===0?c.delete(t):c.set(t,n)}function $(e){let t=c.get(e.name);if(t){let n=[];t.forEach(o=>{o.Callback(e)&&n.push(o)}),n.length>0&&(t=t.filter(o=>!n.includes(o)),t.length===0?c.delete(e.name):c.set(e.name,t))}}function ee(e,...t){[e,...t].forEach(o=>{c.delete(o)})}function te(){c.clear()}function E(e){Ct(vt,e)}var bt=a(l.Dialog),St=0,Mt=1,Dt=2,Et=3,yt=4,Rt=5,m=new Map;function Lt(){let e;do e=g();while(m.has(e));return e}function ne(e,t,n){let o=m.get(e);o&&(n?o.resolve(JSON.parse(t)):o.resolve(t),m.delete(e))}function oe(e,t){let n=m.get(e);n&&(n.reject(t),m.delete(e))}function h(e,t){return new Promise((n,o)=>{let i=Lt();t=t||{},t["dialog-id"]=i,m.set(i,{resolve:n,reject:o}),bt(e,t).catch(r=>{o(r),m.delete(i)})})}function ie(e){return h(St,e)}function re(e){return h(Mt,e)}function le(e){return h(Dt,e)}function x(e){return h(Et,e)}function ae(e){return h(yt,e)}function se(e){return h(Rt,e)}var Ot=a(l.ContextMenu),kt=0;function Nt(e,t,n,o){Ot(kt,{id:e,x:t,y:n,data:o})}function ue(){window.addEventListener("contextmenu",It)}function It(e){let t=e.target,n=window.getComputedStyle(t).getPropertyValue("--custom-contextmenu");if(n=n?n.trim():"",n){e.preventDefault();let o=window.getComputedStyle(t).getPropertyValue("--custom-contextmenu-data");Nt(n,e.clientX,e.clientY,o);return}zt(e)}function zt(e){let t=e.target;switch(window.getComputedStyle(t).getPropertyValue("--default-contextmenu").trim()){case"show":return;case"hide":e.preventDefault();return;default:if(t.isContentEditable)return;let i=window.getSelection(),r=i.toString().length>0;if(r)for(let u=0;u<i.rangeCount;u++){let s=i.getRangeAt(u).getClientRects();for(let w=0;w<s.length;w++){let C=s[w];if(document.elementFromPoint(C.left,C.top)===t)return}}if((t.tagName==="INPUT"||t.tagName==="TEXTAREA")&&(r||!t.readOnly&&!t.disabled))return;e.preventDefault()}}function ce(e,t=null){let n=new M(e,t);E(n)}function At(){document.querySelectorAll("[wml-event]").forEach(function(t){let n=t.getAttribute("wml-event"),o=t.getAttribute("wml-confirm"),i=t.getAttribute("wml-trigger")||"click",r=function(){if(o){x({Title:"Confirm",Message:o,Detached:!1,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(u){u!=="No"&&ce(n)});return}ce(n)};t.removeEventListener(i,r),t.addEventListener(i,r)})}function de(e){wails.Window[e]===void 0&&console.log("Window method "+e+" not found"),wails.Window[e]()}function Tt(){document.querySelectorAll("[wml-window]").forEach(function(t){let n=t.getAttribute("wml-window"),o=t.getAttribute("wml-confirm"),i=t.getAttribute("wml-trigger")||"click",r=function(){if(o){x({Title:"Confirm",Message:o,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(u){u!=="No"&&de(n)});return}de(n)};t.removeEventListener(i,r),t.addEventListener(i,r)})}function Pt(){document.querySelectorAll("[wml-openurl]").forEach(function(t){let n=t.getAttribute("wml-openurl"),o=t.getAttribute("wml-confirm"),i=t.getAttribute("wml-trigger")||"click",r=function(){if(o){x({Title:"Confirm",Message:o,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(u){u!=="No"&&wails.Browser.OpenURL(n)});return}wails.Browser.OpenURL(n)};t.removeEventListener(i,r),t.addEventListener(i,r)})}function j(){At(),Tt(),Pt()}var H=function(e){chrome.webview.postMessage(e)};var fe=new Map;function me(e){let t=new Map;for(let[n,o]of Object.entries(e))typeof o=="object"&&o!==null?t.set(n,me(o)):t.set(n,o);return t}fetch("/wails/flags").then(e=>{e.json().then(t=>{fe=me(t)})});function Bt(e){let t=e.split("."),n=fe;for(let o of t)if(n instanceof Map?n=n.get(o):n=n[o],n===void 0)break;return n}function y(e){return Bt(e)}var W=!1;function jt(e){let t=window.getComputedStyle(e.target).getPropertyValue("--webkit-app-region");return t&&(t=t.trim()),t!=="drag"||e.buttons!==1?!1:e.detail===1}function pe(){window.addEventListener("mousedown",Ut),window.addEventListener("mousemove",Gt),window.addEventListener("mouseup",Ft)}var R=null,we=!1;function ge(e){we=e}function Ht(e){return R?(H("resize:"+R),!0):!1}function Ut(e){if(!Ht())if(jt(e)){if(e.offsetX>e.target.clientWidth||e.offsetY>e.target.clientHeight)return;W=!0}else W=!1}function Ft(e){(e.buttons!==void 0?e.buttons:e.which)>0&&U()}function U(){document.body.style.cursor="default",W=!1}function d(e){document.documentElement.style.cursor=e||Zt,R=e}function Gt(e){if(W){W=!1,(e.buttons!==void 0?e.buttons:e.which)>0&&H("drag");return}we&&Yt(e)}var Zt="auto";function Yt(e){let t=y("system.resizeHandleHeight")||5,n=y("system.resizeHandleWidth")||5,o=y("resizeCornerExtra")||10,i=window.outerWidth-e.clientX<n,r=e.clientX<n,u=e.clientY<t,p=window.outerHeight-e.clientY<t,s=window.outerWidth-e.clientX<n+o,w=e.clientX<n+o,C=e.clientY<t+o,F=window.outerHeight-e.clientY<t+o;!r&&!i&&!u&&!p&&R!==void 0?d():s&&F?d("se-resize"):w&&F?d("sw-resize"):w&&C?d("nw-resize"):C&&s?d("ne-resize"):r?d("w-resize"):u?d("n-resize"):p?d("s-resize"):i&&d("e-resize")}window.wails={...he(null),Capabilities:{},clientId:L};fetch("/wails/capabilities").then(e=>{e.json().then(t=>{window.wails.Capabilities=t})});window._wails={dialogCallback:ne,dialogErrorCallback:oe,dispatchWailsEvent:$,callCallback:Z,callErrorCallback:Y,endDrag:U,setResizable:ge};function he(e){return{Clipboard:{...O},Application:{...I,GetWindowByName(t){return he(t)}},System:T,Screens:A,Browser:P,Call:Q,CallByID:V,CallByName:X,Plugin:q,WML:{Reload:j},Dialog:{Info:ie,Warning:re,Error:le,Question:x,OpenFile:ae,SaveFile:se},Events:{Emit:E,On:_,Once:K,OnMultiple:D,Off:ee,OffAll:te},Window:J(e)}}ue();pe();document.addEventListener("DOMContentLoaded",function(){j()});})();
