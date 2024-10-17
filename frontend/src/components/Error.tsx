// create me an Error component styled with TailwindCSS
// I should display an error message

import { useNavigate } from "react-router-dom";

// I should have a retry button
const Error = () => {
    const navigate = useNavigate();


  return (
    <div className="flex flex-col items-center justify-center h-screen">
        <div className="flex flex-col items-center justify-center w-2/5 h-2/5 border-2 rounded-md shadow-slate-400 border-red-300">            
            <h1 className="text-3xl font-bold text-red-500">An Error occured.</h1>
            <h1 className="text-3xl font-bold text-red-500">Please navigate to main page.</h1>
            <button className="mt-4 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-400 focus:outline-none focus:ring-2 focus:ring-red-600 focus:ring-opacity-50"
            onClick={()=> navigate('/')}
            >Home page</button>
        </div>
    </div>
  );
};

export default Error;