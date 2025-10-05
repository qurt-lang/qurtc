const go = new Go();
WebAssembly.instantiateStreaming(fetch("/assets/wasm/qurtc.wasm"), go.importObject).then((result) => {
	go.run(result.instance);
});

const lessons = [
	{
		title: 'Қош келдіңіз!',
		content: `
<h1>Қош келдіңіз!</h1>
<p>Qurt - қазақ тілінде бағдарламалау тілі. Бұл тур сізге бағдарламалаудың негіздерін үйретеді.</p>

<h2>Бірінші бағдарлама</h2>
<p>Әрбір бағдарлама <code>негізгі</code> функциясынан басталады. Бұл - бағдарламаның бастапқы нүктесі.</p>

<p>Оң жақтағы редакторда қарапайым бағдарлама бар. "Іске қосу" батырмасын басып көріңіз!</p>

<ul>
<li><code>жаз()</code> - экранға мәтін шығарады</li>
<li>Мәтінді тырнақшада жазамыз</li>
<li>Әр команда жаңа жолдан басталады</li>
</ul>
`,
		code: 'функция ештеңе негізгі() {\n    жаз("Сәлем, Qurt!");\n    жаз("Менің алғашқы бағдарламам");\n}'
	},
	{
		title: 'Айнымалылар',
		content: `
<h1>Айнымалылар</h1>
<p>Айнымалылар - деректерді сақтайтын "қорапшалар". Оларға атау беріп, мән сақтаймыз.</p>

<h2>Айнымалы жасау</h2>
<p>Айнымалы жасау үшін <code>айнымалы</code> сөзін қолданамыз:</p>

<ul>
<li><code>айнымалы аты жол</code> - мәтін сақтайды</li>
<li><code>айнымалы жасы бүтін</code> - бүтін сан сақтайды</li>
<li><code>=</code> белгісімен мән береміз</li>
</ul>

<p>Айнымалыны жасап, оны экранға шығарып көріңіз.</p>
`,
		code: 'функция ештеңе негізгі() {\n    айнымалы аты жол = "Айгүл";\n    айнымалы жасы бүтін = 25;\n    \n    жаз("Менің атым: ");\n    жаз(аты);\n    жаз("Менің жасым: ");\n    жаз(жасы);\n}'
	},
	{
		title: 'Сандармен жұмыс',
		content: `
<h1>Сандармен жұмыс</h1>
<p>Сандарға математикалық амалдар қолдана аламыз.</p>

<h2>Математикалық амалдар</h2>
<ul>
<li><code>+</code> - қосу</li>
<li><code>-</code> - алу</li>
<li><code>*</code> - көбейту</li>
<li><code>/</code> - бөлу</li>
<li><code>%</code> - қалдық табу</li>
</ul>

<p>Санды айнымалыға сақтап, онымен әртүрлі амалдар жасай аламыз.</p>
`,
		code: 'функция ештеңе негізгі() {\n    айнымалы а бүтін = 15;\n    айнымалы б бүтін = 3;\n    \n    айнымалы қосынды бүтін = а + б;\n    айнымалы көбейтінді бүтін = а * б;\n    \n    жаз("Қосынды: ");\n    жаз(қосынды);\n    жаз("Көбейтінді: ");\n    жаз(көбейтінді);\n}'
	},
	{
		title: 'Шарттар',
		content: `
<h1>Шарттар</h1>
<p>Шарттар арқылы бағдарлама әртүрлі жағдайларда әртүрлі әрекет жасайды.</p>

<h2>Егер командасы</h2>
<p><code>егер</code> - шарт дұрыс болса код орындалады:</p>

<ul>
<li><code>==</code> - тең</li>
<li><code>></code> - артық</li>
<li><code>&lt;</code> - кем</li>
<li><code>>=</code> - артық немесе тең</li>
</ul>

<p><code>әйтпесе</code> - шарт жалған болса орындалады.</p>
`,
		code: 'функция ештеңе негізгі() {\n    айнымалы жасы бүтін = 20;\n    \n    егер(жасы >= 18) {\n        жаз("Сіз ересексіз");\n    } әйтпесе {\n        жаз("Сіз кәмелетсізсіз");\n    }\n}'
	},
	{
		title: 'Циклдер',
		content: `
<h1>Циклдер</h1>
<p>Циклдер кодты бірнеше рет қайталауға көмектеседі.</p>

<h2>Қайтала циклі</h2>
<p><code>қайтала</code> командасы кодты белгілі санда қайталайды:</p>

<ul>
<li>Басталу мәнін көрсетеміз</li>
<li>Аяқталу шартын жазамыз</li>
<li>Қадамды анықтаймыз</li>
</ul>

<p>Мысалда санағыш 1-ден 5-ке дейін санайды.</p>
`,
		code: 'функция ештеңе негізгі() {\n    жаз("Санақ басталды:");\n    \n    қайтала(айнымалы i бүтін = 1; i <= 5; i = i + 1) {\n        жаз(i);\n    }\n    \n    жаз("Санақ аяқталды!");\n}'
	},
	{
		title: 'Функциялар',
		content: `
<h1>Функциялар</h1>
<p>Функция - кодтың бір бөлігін атау беріп, оны қайта-қайта қолдана аламыз.</p>

<h2>Функция жасау</h2>
<p><code>функция</code> сөзімен функция жасаймыз:</p>

<ul>
<li>Функцияға атау береміз</li>
<li>Қайтаратын түрін көрсетеміз</li>
<li><code>ештеңе</code> - ештеңе қайтармайды</li>
<li>Функцияны атымен шақырамыз</li>
</ul>
`,
		code: 'функция ештеңе сәлемДе() {\n    жаз("Сәлеметсіз бе!");\n    жаз("Қош келдіңіз!");\n}\n\nфункция ештеңе негізгі() {\n    сәлемДе();\n    сәлемДе();\n}'
	},
	{
		title: 'Параметрлер мен қайтару',
		content: `
<h1>Параметрлер мен қайтару</h1>
<p>Функцияларға мән беріп, нәтиже алуға болады.</p>

<h2>Параметрлер</h2>
<ul>
<li>Функцияға жақшада параметр береміз</li>
<li><code>қайтар</code> - нәтиже қайтарады</li>
<li>Функция нәтижесін айнымалыға сақтаймыз</li>
</ul>

<h2>Құттықтаймыз!</h2>
<p>Сіз Qurt тілінің негізгі ұғымдарын үйрендіңіз! Енді өз бағдарламаларыңызды жаза аласыз.</p>
`,
		code: 'функция бүтін қос(а бүтін, б бүтін) {\n    айнымалы нәтиже бүтін = а + б;\n    қайтар нәтиже;\n}\n\nфункция ештеңе негізгі() {\n    айнымалы сан1 бүтін = 10;\n    айнымалы сан2 бүтін = 25;\n    \n    айнымалы қосынды бүтін = қос(сан1, сан2);\n    \n    жаз("Қосынды: ");\n    жаз(қосынды);\n}'
	}
];

let currentLesson = 0;
const totalLessons = lessons.length;
const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');
const runBtn = document.getElementById('runBtn');
const codeEditor = document.getElementById('codeEditor');
const output = document.getElementById('output');
const lessonContent = document.getElementById('lessonContent');
const currentPageSpan = document.getElementById('currentPage');
const totalPagesSpan = document.getElementById('totalPages');

totalPagesSpan.textContent = totalLessons;

function updateLesson() {
	const lesson = lessons[currentLesson];
	lessonContent.innerHTML = lesson.content;
	codeEditor.value = lesson.code;
	prevBtn.disabled = currentLesson === 0;
	nextBtn.disabled = currentLesson === totalLessons - 1;
	if (currentLesson === totalLessons - 1) {
		nextBtn.textContent = 'Аяқтау ✓';
	} else {
		nextBtn.textContent = 'Алға →';
	}
	currentPageSpan.textContent = currentLesson + 1;
	output.textContent = '';
}
prevBtn.addEventListener('click', () => {
	if (currentLesson > 0) {
		currentLesson--;
		updateLesson();
	}
});
nextBtn.addEventListener('click', () => {
	if (currentLesson < totalLessons - 1) {
		currentLesson++;
		updateLesson();
	} else {
		alert('Құттықтаймыз! Сіз турды аяқтадыңыз!');
	}
});
runBtn.addEventListener('click', () => {
	output.textContent = qurtExec(codeEditor.value);
});
updateLesson();
