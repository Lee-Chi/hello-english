(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[909],{2377:function(e,t,a){Promise.resolve().then(a.bind(a,6082))},9896:function(e,t,a){"use strict";var s=a(7908);let r="https://hello-english-1738.de.r.appspot.com",n=async e=>{let t="".concat(r,"/").concat(e);console.log(t);try{return(await s.Z.get(t)).data}catch(e){console.error(e)}},l=async(e,t)=>{let a="".concat(r,"/").concat(e);console.log(a);try{return(await s.Z.post(a,t)).data}catch(e){console.error(e)}};t.Z={get:n,post:l}},135:function(e,t,a){"use strict";var s=a(9896);let r=async e=>{let t=await s.Z.post("api/sentence/translate",{sentence:e});return t.errorCode>0?(console.error(t),""):t.translation},n=async e=>{let t=await s.Z.post("api/sentence/check",{sentence:e});return t.errorCode>0?(console.error(t),""):t.answer},l=async e=>{let t=await s.Z.post("api/sentence/advise",{sentence:e});return t.errorCode>0?(console.error(t),{}):(console.log(t),t.advice)},o=async e=>{let t=await s.Z.post("api/sentence/practice/ready",{topic:e});return t.errorCode>0?(console.error(t),{}):t.question},i=async(e,t)=>{let a=await s.Z.post("api/sentence/practice/submit",{question:e,answer:t});return a.errorCode>0?(console.error(a),{}):(console.log(a),{correctAnswer:a.correctAnswer,advices:a.advices})};t.Z={translate:r,check:n,advise:l,readyPractice:o,submitPractice:i}},9124:function(e,t,a){"use strict";a.d(t,{g:function(){return o}});var s=a(3827),r=a(4090);class n extends r.Component{render(){let{className:e,children:t}=this.props;return(0,s.jsx)("div",{className:e,children:(0,s.jsxs)("div",{className:"w-[300px] min-h-screen bg-cerulean-400",children:[(0,s.jsxs)("div",{className:"w-full h-[100px] bg-aero-100 flex border-carrot_orange-500",children:[(0,s.jsx)("div",{className:"pt-2 px-2",children:(0,s.jsx)("div",{className:"w-20 h-20 rounded-full bg-aero-700 p-1",children:(0,s.jsx)("img",{className:"w-full h-full rounded-full",src:"https://picsum.photos/300/300?people=10",alt:"圖片"})})}),(0,s.jsxs)("div",{className:"pt-4",children:[(0,s.jsx)("h1",{className:"text-xl font-bold text-carrot_orange-500",children:"Hello English"}),(0,s.jsx)("p",{className:"text-xs text-carrot_orange-600 mt-1",children:"Improve on your English every day!"})]})]}),(0,s.jsx)("div",{className:"w-[300px] h-[300px] grid items-center justify-center",children:(0,s.jsx)("div",{className:"mt-10 p-2 rounded-sm bg-cerulean-100 border border-cerulean-600 shadow-md shadow-cerulean-600",children:(0,s.jsx)("div",{className:"p-14 rounded-sm bg-cerulean-200 border border-cerulean-600 shadow-md shadow-cerulean-600",children:t})})})]})})}}var l=a(7907);function o(e){let{className:t,children:a}=e,r=(0,l.useRouter)();return(0,s.jsx)(n,{className:t,children:(0,s.jsxs)("ul",{className:"border-carrot_orange-600 text-carrot_orange-600",children:[(0,s.jsx)("li",{className:"border-b border-carrot_orange-600",onClick:()=>r.push("/word"),children:(0,s.jsx)("span",{className:"hover:text-carrot_orange-800 hover:cursor-pointer",children:"WORD"})}),(0,s.jsx)("li",{className:"pt-10 border-b border-carrot_orange-600",onClick:()=>r.push("/sentence"),children:(0,s.jsx)("span",{className:"hover:text-carrot_orange-800 hover:cursor-pointer",children:"SENTENCE"})}),(0,s.jsx)("li",{className:"pt-10 border-b border-carrot_orange-600",children:(0,s.jsx)("span",{className:"hover:text-carrot_orange-800 hover:cursor-not-allowed",children:"PARAGRAPH"})})]})})}},6082:function(e,t,a){"use strict";a.r(t),a.d(t,{default:function(){return g}});var s=a(3827),r=a(4090),n=a(9124),l=a(135),o=a(243);let i=r.forwardRef(function(e,t){let{title:a,titleId:s,...n}=e;return r.createElement("svg",Object.assign({xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24",strokeWidth:1.5,stroke:"currentColor","aria-hidden":"true","data-slot":"icon",ref:t,"aria-labelledby":s},n),a?r.createElement("title",{id:s},a):null,r.createElement("path",{strokeLinecap:"round",strokeLinejoin:"round",d:"m19.5 8.25-7.5 7.5-7.5-7.5"}))});class c extends r.Component{render(){return(0,s.jsxs)("div",{className:"w-full h-full grid",children:[(0,s.jsx)("div",{className:"",children:(0,s.jsx)("input",{type:"textarea",className:"w-full h-20 p-4 bg-white text-areo text-2xl font-bold rounded-t-md",value:this.state.sentenceWant,onChange:e=>this.handleSentenceWantChange(e.target.value),placeholder:"Enter a sentence..."})}),(0,s.jsxs)("div",{className:"flex h-[640px]",children:[(0,s.jsxs)("div",{className:"w-1/3 h-full bg-columbia_blue-500 hover:border-4 hover:border-aero-400 rounded-bl-md ",children:[(0,s.jsx)("button",{className:"w-full h-[60px] opacity-85 bg-carrot_orange-500 text-yellow_green-200 text-2xl font-bold hover:bg-carrot_orange-400 hover:text-yellow_green-500",onClick:async()=>{if(""===this.state.sentenceWant)return;let e={loading:!0,answer:""};this.handleCheckChange(e);let t=await l.Z.check(this.state.sentenceWant);e.loading=!1,e.answer=t,this.handleCheckChange(e)},children:this.state.check.loading?(0,s.jsxs)("svg",{"aria-hidden":"true",class:"inline w-8 h-8 text-carrot_orange-500 animate-spin dark:text-gray-600 fill-carrot_orange-300",viewBox:"0 0 100 101",fill:"none",xmlns:"http://www.w3.org/2000/svg",children:[(0,s.jsx)("path",{d:"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z",fill:"currentColor"}),(0,s.jsx)("path",{d:"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z",fill:"currentFill"})]}):"Check"}),(0,s.jsx)("p",{className:"break-words p-4 text-gray-800 text-xl",children:this.state.check.answer})]}),(0,s.jsxs)("div",{className:"w-1/3 h-full bg-columbia_blue-400 hover:border-4 hover:border-aero-400",children:[(0,s.jsx)("button",{className:"block w-full h-[60px] opacity-85 bg-carrot_orange-500 text-yellow_green-200 text-2xl font-bold hover:bg-carrot_orange-400 hover:text-yellow_green-500",onClick:async()=>{if(""===this.state.sentenceWant)return;let e={loading:!0,translation:""};this.handleTranslateChange(e);let t=await l.Z.translate(this.state.sentenceWant);e.loading=!1,e.translation=t,this.handleTranslateChange(e)},children:this.state.translate.loading?(0,s.jsxs)("svg",{"aria-hidden":"true",class:"inline w-8 h-8 text-carrot_orange-500 animate-spin dark:text-gray-600 fill-carrot_orange-300",viewBox:"0 0 100 101",fill:"none",xmlns:"http://www.w3.org/2000/svg",children:[(0,s.jsx)("path",{d:"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z",fill:"currentColor"}),(0,s.jsx)("path",{d:"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z",fill:"currentFill"})]}):"Translate"}),(0,s.jsx)("p",{className:"break-words p-4 text-gray-800 text-xl",children:this.state.translate.translation})]}),(0,s.jsxs)("div",{className:"w-1/3 h-full bg-columbia_blue-300 hover:border-4 hover:border-aero-400 rounded-br-md",children:[(0,s.jsx)("button",{className:"block w-full h-[60px] opacity-85 bg-carrot_orange-500 text-yellow_green-200 text-2xl font-bold hover:bg-carrot_orange-400 hover:text-yellow_green-500",onClick:async()=>{if(""===this.state.sentenceWant)return;let e={loading:!0,revised:"",reasons:[]};this.setState({advise:e});let t=await l.Z.advise(this.state.sentenceWant);e.loading=!1,e.revised=t.revised,e.reasons=t.reasons,this.setState({advise:e})},children:this.state.advise.loading?(0,s.jsxs)("svg",{"aria-hidden":"true",class:"inline w-8 h-8 text-carrot_orange-500 animate-spin dark:text-gray-600 fill-carrot_orange-300",viewBox:"0 0 100 101",fill:"none",xmlns:"http://www.w3.org/2000/svg",children:[(0,s.jsx)("path",{d:"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z",fill:"currentColor"}),(0,s.jsx)("path",{d:"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z",fill:"currentFill"})]}):"Advise"}),(0,s.jsx)("div",{className:"",children:this.state.advise&&(0,s.jsxs)("div",{children:[(0,s.jsx)("p",{className:"break-words p-4 text-gray-800 text-xl",children:this.state.advise.revised}),(0,s.jsx)("ul",{children:this.state.advise.reasons.map((e,t)=>(0,s.jsxs)("li",{className:"text-gray-800 text-sm",children:[(0,s.jsxs)("span",{className:"font-bold",children:["[",e.type,"]"]}),(0,s.jsx)("span",{className:"ml-1",children:e.message})]},t))})]})})]})]})]})}constructor(e){super(e),this.handleSentenceWantChange=e=>{this.setState({sentenceWant:e})},this.handleCheckChange=async e=>{this.setState({check:e})},this.handleTranslateChange=async e=>{this.setState({translate:e})},this.state={sentenceWant:"",check:{loading:!1,answer:""},translate:{loading:!1,translation:""},advise:{loading:!1,revised:"",reasons:[]}}}}let d=["Program","Travel","Food","Health","Science","Technology","Education","Business","Art","Music","Movie","Book","History","Politics","Social","Environment","Weather","Life","Love","Family","Friend","Work","Study","Game","Sport"],h={loading:!1,question:"",answer:""},u={loading:!1,correctAnswer:"",advices:[]};class x extends r.Component{render(){return(0,s.jsxs)("div",{className:"w-full h-full grid gap-8 font-mono",children:[(0,s.jsxs)("div",{className:"flex w-full gap-2",children:[(0,s.jsxs)("div",{className:"flex",children:[(0,s.jsx)(o.Input,{type:"text",value:this.state.topic.value,disabled:!0,className:"rounded-r-none !font-mono"}),(0,s.jsxs)(o.Menu,{open:this.state.topic.openMenu,handler:this.handleTopicMenuChange,className:"",children:[(0,s.jsx)(o.MenuHandler,{children:(0,s.jsx)(o.Button,{size:"sm",label:"Topic",className:"rounded-l-none text-white bg-blue-gray-700 hover:bg-blue-gray-500 focus:bg-blue-gray-500",children:(0,s.jsx)(i,{className:"ml-1 w-4 h-4"})})}),(0,s.jsx)(o.MenuList,{className:"max-h-72 font-mono",children:d.map((e,t)=>(0,s.jsx)(o.MenuItem,{onClick:()=>this.handleTopicValueChange(e),children:e},t))})]})]}),(0,s.jsx)(o.Button,{size:"sm",className:"bg-red-600",loading:this.state.ready.loading,onClick:async()=>{if(""===this.state.topic.value)return;let e=await l.Z.readyPractice(this.state.topic.value);this.handleReadyQuestionChange(e),this.handleReadyAnswerChange(""),this.handleSubmitChange(u)},children:"Go"})]}),(0,s.jsxs)("div",{className:"grid",children:[(0,s.jsx)("h1",{className:"block text-base font-bold text-carrot_orange-500",children:"Question:"}),(0,s.jsx)("div",{className:"h-20 block  text-carrot_orange-600",children:(0,s.jsx)("p",{className:"p-2 break-words",children:this.state.ready.question})})]}),(0,s.jsxs)("div",{className:"grid",children:[(0,s.jsx)("h1",{className:"block text-base font-bold text-carrot_orange-500",children:"Your Answer:"}),(0,s.jsx)("textarea",{className:"block w-full resize-none rounded-md bg-gray-800 text-carrot_orange-600 border-0 focus:outline-0 focus:ring-0 focus:border-0",value:this.state.ready.answer,onChange:e=>this.handleReadyAnswerChange(e.target.value),rows:3}),(0,s.jsx)("div",{className:"flex justify-end bg-gray-800 rounded-b-md p-2",children:(0,s.jsx)(o.Button,{size:"sm",className:"bg-red-600",loading:this.state.submit.loading,onClick:async()=>{if(""===this.state.ready.question)return;this.handleSubmitChange({loading:!0,correctAnswer:"",advices:[]});let e=await l.Z.submitPractice(this.state.ready.question,this.state.ready.answer);this.handleSubmitChange({loading:!1,correctAnswer:e.correctAnswer,advices:e.advices})},children:"Submit"})})]}),(0,s.jsxs)("div",{className:"grid",children:[(0,s.jsx)("h1",{className:"block text-base font-bold text-carrot_orange-500",children:"Advices:"}),(0,s.jsxs)("div",{className:"grid gap-1 p-2",children:[(0,s.jsx)("p",{className:"p-2 block break-words text-carrot_orange-600",children:this.state.submit.correctAnswer}),(0,s.jsx)(o.List,{children:this.state.submit.advices.map((e,t)=>(0,s.jsx)(o.ListItem,{className:"p-2 block break-words italic text-carrot_orange-600 hover:bg-gray-600 hover:text-carrot_orange-700 active:bg-gray-600 active:text-carrot_orange-700 focus:bg-gray-600 focus:text-carrot_orange-700",ripple:!1,children:e},t))})]})]})]})}constructor(e){super(e),this.handleTopicMenuChange=e=>{let t=this.state.topic;t.openMenu=e,this.setState({topic:t})},this.handleTopicValueChange=e=>{let t=this.state.topic;t.value=e,this.setState({topic:t})},this.handleReadyQuestionChange=e=>{let t=this.state.ready;t.question=e,this.setState({ready:t})},this.handleReadyAnswerChange=e=>{let t=this.state.ready;t.answer=e,this.setState({ready:t})},this.handleSubmitChange=e=>{this.setState({submit:e})},this.state={topic:{openMenu:!1,value:d[0]},ready:h,submit:u}}}class g extends r.Component{render(){let e=[{label:"Tunning",value:"tunning",content:(0,s.jsx)(c,{})},{label:"Practice",value:"practice",content:(0,s.jsx)(x,{})}];return(0,s.jsxs)("div",{className:"w-full min-h-screen flex bg-columbia_blue font-mono",children:[(0,s.jsx)(n.g,{}),(0,s.jsx)("div",{className:"w-full min-h-screen bg-cerulean-900 grid justify-center content-center",children:(0,s.jsxs)(o.Tabs,{value:"tunning",children:[(0,s.jsx)(o.TabsHeader,{className:"bg-blue-gray-200 mb-1",children:e.map(e=>(0,s.jsx)(o.Tab,{label:e.label,value:e.value,className:"font-mono text-xl font-bold text-blue-gray-600",children:e.label},e.value))}),(0,s.jsx)(o.TabsBody,{className:"w-[1024px] h-[760px] rounded-md bg-blue-gray-800",children:e.map(e=>(0,s.jsx)(o.TabPanel,{value:e.value,className:"font-mono",children:e.content},e.value))})]})})]})}}}},function(e){e.O(0,[384,823,971,69,744],function(){return e(e.s=2377)}),_N_E=e.O()}]);