import { Match } from "../services/MatchService"

interface MatchElementProps {
    match : Match
    onClickMatch : (chatId: string, match: Match) => void
    selectedMatch?: Match | null
}


const MatchElement = ({match, onClickMatch, selectedMatch} : MatchElementProps) => {
  return (
    <div className={`flex flex-row justify-center items-center border-y rounded-md py-3 hover:cursor-pointer hover:bg-gray-700 ${selectedMatch?.id === match.id ? 'bg-gray-700' : ''} w-full`}
        onClick={()=>onClickMatch(match.chatId, match)}
    >
        <div className="flex flex-row justify-center items-center w-full lg:w-52">
            <p className="desc">{match.team1}</p>
        </div>
        {/* a small div to display the score */}
        <div className="flex flex-row justify-center items-center w-full lg:w-10">
            <p className="desc">{match.score1}</p>
        </div>
        <div className="flex flex-col justify-center items-center w-full lg:w-[34rem]">
            <div className="flex flex-row justify-center items-center gap-5">
                <img src={match.tournament_icon} alt="Tournament icon" className="w-3/12 lg:w-1/12" />
                <p className="orange_gradient">{match.tournament_name}</p>
            </div>
        </div>
        <div className="flex flex-row justify-center items-center w-full lg:w-10">
            <p className="desc">{match.score2}</p>
        </div>
        <div className="flex flex-row justify-center items-center w-full lg:w-52">
            <p className="desc">{match.team2}</p>
        </div>
    </div>
  )
}

export default MatchElement