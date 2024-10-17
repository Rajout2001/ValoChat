// import { useState } from "react";
import { Link, useNavigate } from "react-router-dom"
import { useAuth } from "../context/AuthProvider";

const Feed = () => {
    const {user} = useAuth();
    const navigate = useNavigate();

  return (
    <section className="w-full flex flex-col items-center gap-3">
        <h1 className="text-center">
            <span className="head_text">
                Join us to discuss about
            </span>
            <br/> 
            <span className="orange_gradient text-center lg:text-3xl">Valorant Match Results</span>
        </h1>
        <p className="desc text-center">Our amazing plateform will let you vote your favorite team
        and let you discuss the match's details with passionate people like you!
        </p>

        <div className="flex flex-col justify-center items-center gap-6">
            <p
                className="text-center text-slate-400 text-2xl"
                >Don't waste your time and come enjoy with us..
            </p>
            {user && (
                <div className='text-black bg-slate-400 hover:bg-slate-300 font-bold py-4 px-6 my-2 rounded hover:cursor-pointer' 
                    onClick={() => navigate('/matchs')}
                >
                    <button>
                        Go to Matchs
                    </button>
                </div>
            )}
            <div className='mx-6'>
                {user ? (
                    <div className="desc text-center">
                        Our amazing plateform is 100% free, No credit card required.
                    </div>  
                ) : (
                    <button className="my-5">
                        <Link to='/login' className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-6 rounded'> 
                            Login
                        </Link>
                    </button>
                )

                }
            
            </div>
        </div>
        <div className="my-2 w-full h-full lg:w-2/3 lg:h-1/2 flex justify-center items-center">
            <img
                className="w-1/2 h-1/2 rounded-lg shadow-lg border-2 border-gray-200"
                src="/feed-team.png"
                alt="feed-team"
            />
        </div> 
        
    </section>
  )
}

export default Feed
