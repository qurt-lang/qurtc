import { App } from './App';
import { createRoot } from "react-dom/client";
import "./index.css";

createRoot(document.getElementById('root')!).render(
	<>
		<header className="flex justify-between items-center px-10 py-2 bg-white shadow-md w-full">
			<a href="/">
				<div className="flex items-center">
					<img src="/assets/images/logo.png"
						className="w-[35px] h-[35px] bg-[#ff9c1a] rounded-[5px] mr-[10px]" />
					<span className="font-bold text-[20px] text-[#575e75]">Qurt Қазақша</span>
				</div>
			</a>
		</header>
		<div>
			<App />
		</div>
	</>
);
