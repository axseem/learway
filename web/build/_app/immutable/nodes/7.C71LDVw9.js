import{a as lt}from"../chunks/api.Zt1bz1uW.js";import{a as f,t as D,n as H,s as t,c as e,d as ot}from"../chunks/disclose-version.45ewW4YV.js";import{r as m,p as it,a as nt,s as c,g as r,A as T,c as g}from"../chunks/runtime.BZ58vcTB.js";import{B as x,e as ct}from"../chunks/button.6U0bErQb.js";import{e as A,s as d,a as w}from"../chunks/render.ras3OibP.js";import{p as _}from"../chunks/proxy.CQrFxjn9.js";function b(v,o,a){a?v.classList.add(o):v.classList.remove(o)}const dt=async({fetch:v,params:o})=>({deck:await(await v(lt+"/deck/"+o.id,{method:"get"})).json()}),bt=Object.freeze(Object.defineProperty({__proto__:null,load:dt},Symbol.toStringTag,{value:"Module"}));var vt=H('<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide-icon lucide lucide-arrow-left h-4 w-4"><path d="m12 19-7-7 7-7"></path><path d="M19 12H5"></path></svg>'),ut=H('<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide-icon lucide lucide-arrow-right h-4 w-4"><path d="M5 12h14"></path><path d="m12 5 7 7-7 7"></path></svg>'),pt=D('<li class="bg-card flex w-full divide-x overflow-hidden rounded-lg border p-6"><p class="w-full p-2 pr-8"> </p> <p class="w-full p-2 pl-8"> </p></li>'),ft=D('<div class="flex h-full w-full justify-center"><div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6"><div class="pb-2"><h1 class="text-2xl font-medium"> </h1> <button class="text-muted-foreground font-mono text-sm"> </button></div> <button class="ring-border bg-background flex aspect-[2/1] w-full shrink-0 items-center justify-center rounded-2xl border p-8 ring-inset transition-transform hover:scale-105 active:scale-100"><p class="text-xl"> </p></button> <div class="flex items-center"><div class="w-full flex-col"><p class="font-mono text-sm">front</p> <p class="font-mono text-sm">back</p></div> <div class="flex w-full items-center justify-center gap-4"><!> <div class="font-mono"> </div> <!></div> <div class="flex w-full justify-end"><!></div></div> <details open><summary class="select-none pb-2 pt-4 text-lg font-medium hover:cursor-pointer">List of cards</summary> <ul class="flex flex-col gap-4 pb-12"></ul></details></div></div>');function kt(v,o){nt(o,!0);let a=_(o.data.deck),i=g(0),u=g(!1),l=g(!1);const N=()=>c(l,!r(l)),U=()=>{c(i,(r(i)+1)%a.cards.length),c(l,_(r(u)))},q=()=>{c(i,(r(i)+a.cards.length-1)%a.cards.length),c(l,!1)},E=()=>{c(u,!r(u)),c(l,_(r(u)))},F=s=>{navigator.clipboard.writeText(s).then(()=>{alert("Copied to clipboard")})};var k=ft(),G=e(k),y=e(G),$=e(y),I=e($),j=t(t($,!0)),J=e(j),h=t(t(y,!0)),K=e(h),Q=e(K);m(()=>d(Q,a.cards[r(i)][Number(r(l))]));var B=t(t(h,!0)),P=e(B),z=e(P),R=t(t(z,!0)),C=t(t(P,!0)),L=e(C);x(L,{variant:"outline",size:"icon",class:"h-12 w-12 rounded-full",$$events:{click:q},children:(s,p)=>{var n=vt();f(s,n)}});var M=t(t(L,!0)),V=e(M),W=t(t(M,!0));x(W,{variant:"outline",size:"icon",class:"h-12 w-12 rounded-full",$$events:{click:U},children:(s,p)=>{var n=ut();f(s,n)}});var X=t(t(C,!0)),Y=e(X);x(Y,{variant:"outline",$$events:{click:E},children:(s,p)=>{var n=ot(s);m(()=>d(n,r(u)?"back first":"front first")),f(s,n)}});var Z=t(t(B,!0)),tt=e(Z),et=t(t(tt,!0));ct(et,73,()=>a.cards,(s,p,n)=>{var S=pt(),O=e(S),rt=e(O),at=t(t(O,!0)),st=e(at);m(()=>{d(rt,T(p)[0]),d(st,T(p)[1])}),f(s,S)}),m(()=>{d(I,a.title),d(J,`id:${w(a.id)}`),b(h,"ring-4",r(l)),b(z,"opacity-20",r(l)),b(R,"opacity-20",!r(l)),d(V,`${w(r(i)+1)}/${w(a.cards.length)}`)}),A("click",j,()=>F(window.location.href),!1),A("click",h,N,!1),f(v,k),it()}export{kt as component,bt as universal};