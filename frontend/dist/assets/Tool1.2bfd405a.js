import{d as p,e as _,r as l,o as m,c as B,a as u,w as f,f as C,n as a,_ as v}from"./index.8460c796.js";function x(n){return window.go.main.App.Lifestyle(n)}const w={class:"container"},A=p({__name:"Tool1",setup(n){const s=_({addresslist:"",response:"",disabled:!1}),i=async()=>{s.disabled=!0;try{const o=await x(s.addresslist);if(o.code==200){let e=o.data;s.response=e.resultcontent,d("success",e.resultlen)}else a.error({message:"error",description:"\u8BF7\u6C42\u9519\u8BEF"})}catch{a.error({message:"error",description:"\u8BF7\u6C42\u9519\u8BEF"})}s.disabled=!1},d=(o,e)=>{a[o]({message:`${o}`,description:`\u5171\u8BA1\u68C0\u6D4B\u5230${e}\u4E2A\u5B58\u6D3B\u8D44\u4EA7`})};return(o,e)=>{const r=l("a-textarea"),c=l("a-button");return m(),B("div",w,[u(r,{class:"source",value:s.addresslist,"onUpdate:value":e[0]||(e[0]=t=>s.addresslist=t),placeholder:"\u5F85\u68C0\u6D4B\u8D44\u4EA7(\u8F93\u5165\u5185\u5BB9\uFF0C\u7528\u6362\u884C\u3001\u9017\u53F7\u6216\u7A7A\u683C\u5206\u9694)","auto-size":{minRows:10,maxRows:10}},null,8,["value"]),u(c,{class:"check",type:"primary",onClick:i,disabled:s.disabled},{default:f(()=>e[2]||(e[2]=[C("\u68C0\u6D4B\u8D44\u4EA7")])),_:1},8,["disabled"]),u(r,{class:"response",value:s.response,"onUpdate:value":e[1]||(e[1]=t=>s.response=t),placeholder:"\u8D44\u4EA7\u5B58\u6D3B\u60C5\u51B5","auto-size":{minRows:10,maxRows:10}},null,8,["value"])])}}});const F=v(A,[["__scopeId","data-v-d2ee14a5"]]);export{F as default};