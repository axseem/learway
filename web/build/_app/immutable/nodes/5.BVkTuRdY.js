import{g as T}from"../chunks/entry.CnPrUYvS.js";import{a as d,t as i,s as e,c as o}from"../chunks/disclose-version.CfobuPx6.js";import{p as I,a as A,D as s}from"../chunks/runtime.r5-_hA6s.js";import{e as D,B as h}from"../chunks/button.DYXPdrKJ.js";import{b as J}from"../chunks/render.Cxa2e9Dg.js";import{r as w}from"../chunks/attributes.Dbg_PeuU.js";import{I as M,b as y}from"../chunks/input.D5xQT3G4.js";import{p as N}from"../chunks/proxy.BnQUgK9D.js";import{a as U}from"../chunks/api.Zt1bz1uW.js";const V=async({parent:n})=>{const{authorized:c}=await n();c==!1&&T("/login")},ee=Object.freeze(Object.defineProperty({__proto__:null,load:V},Symbol.toStringTag,{value:"Module"}));var q=i('<div class="flex flex-col gap-1.5"><p class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"> </p> <li class="flex w-full"><textarea placeholder="Content for the front..." class="bg-card w-full resize-none rounded-l-lg border p-8"></textarea> <textarea placeholder="Content for the back..." class="bg-card w-full resize-none rounded-r-lg border-y border-r p-8"></textarea></li></div>'),E=i("Add card",1),F=i("Create",1),G=i('<div class="flex h-full w-full justify-center"><div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6"><div class="pb-2"><h1 class="text-2xl font-medium">Create new deck</h1></div> <div class="flex flex-col gap-1.5"><label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="title">Title</label> <!></div> <!> <div class="flex place-content-between"><!> <!></div></div></div>');function ae(n,c){A(c,!0);async function $(r){await fetch(U+"/deck",{method:"POST",credentials:"include",body:JSON.stringify(r),headers:{"Content-Type":"application/json"}}).then(l=>{l.status>=400&&l.status<500?l.json().then(a=>alert(a.message)):l.status>=200&&l.status<300&&l.json().then(a=>window.location.replace("/deck/"+a.id))})}let t=N({title:"",cards:[["",""],["",""]]});var v=G(),k=o(v),j=o(k),f=e(e(j,!0)),C=o(f),z=e(e(C,!0));M(z,{id:"title",name:"title",get value(){return t.title},set value(r){t.title=r}});var m=e(e(f,!0));D(m,65,()=>t.cards,(r,l,a)=>{var x=q(),_=o(x),P=o(_);P.nodeValue=`#${J(s(a)+1)}`;var S=e(e(_,!0)),p=o(S);w(p);var g=e(e(p,!0));w(g),y(p,()=>t.cards[s(a)][0],u=>t.cards[s(a)][0]=u),y(g,()=>t.cards[s(a)][1],u=>t.cards[s(a)][1]=u),d(r,x)});var B=e(e(m,!0)),b=o(B);h(b,{variant:"secondary",$$events:{click:()=>t.cards.push(["",""])},children:(r,l)=>{var a=E();d(r,a)}});var O=e(e(b,!0));h(O,{$$events:{click:()=>$(t)},children:(r,l)=>{var a=F();d(r,a)}}),d(n,v),I()}export{ae as component,ee as universal};
