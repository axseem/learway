import{a as p,t as u,c as s}from"../chunks/disclose-version.DjcaTDsi.js";import{s as a,p as f,a as v,g as t,c as h}from"../chunks/runtime.BZ58vcTB.js";import{e as k}from"../chunks/render.FyCEabsq.js";import{r as g}from"../chunks/attributes.CgUNRT15.js";import{b as x}from"../chunks/input.CUIQP067.js";var _=u('<div class="flex h-full w-full justify-center"><div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6"><div class="flex items-center gap-2"><input type="checkbox" name="dark" id="dark"> <label for="dark">Dark mode</label></div></div></div>');function y(d,l){v(l,!0);let e=h(!0);function m(){a(e,!t(e)),localStorage.setItem("theme",t(e)?"dark":"light"),t(e)?document.documentElement.classList.add("dark"):document.documentElement.classList.remove("dark")}localStorage.theme==="dark"||!("theme"in localStorage)&&window.matchMedia("(prefers-color-scheme: dark)").matches?(document.documentElement.classList.add("dark"),a(e,!0)):(document.documentElement.classList.remove("dark"),a(e,!1));var o=_(),c=s(o),i=s(c),r=s(i);g(r),k("change",r,m,!1),x(r,()=>t(e),n=>a(e,n)),p(d,o),f()}export{y as component};