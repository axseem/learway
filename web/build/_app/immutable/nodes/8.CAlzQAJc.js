import{g as m}from"../chunks/entry.CnPrUYvS.js";import{a,t as i,c as r}from"../chunks/disclose-version.CfobuPx6.js";import{p as u,a as f}from"../chunks/runtime.r5-_hA6s.js";import{i as d}from"../chunks/lifecycle.CLwrvhWQ.js";import{B as v}from"../chunks/button.DYXPdrKJ.js";import"../chunks/stores.BZ1xXB1E.js";const g=async({parent:e})=>{const{authorized:o}=await e();o==!1&&m("/login")},B=Object.freeze(Object.defineProperty({__proto__:null,load:g},Symbol.toStringTag,{value:"Module"}));var _=i("Log out",1),h=i('<div class="flex h-full w-full justify-center"><div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6"><!></div></div>');function O(e,o){f(o,!1);const l=()=>{document.cookie="sessionID=; max-age=0; path=/; secure=true;",localStorage.removeItem("username"),window.location.href="/"};d();var t=h(),s=r(t),n=r(s);v(n,{variant:"destructive","data-sveltekit-reload":!0,$$events:{click:l},children:(c,w)=>{var p=_();a(c,p)}}),a(e,t),u()}export{O as component,B as universal};
