import axios from "axios"

export interface Match {
    id: string
    team1: string
    team2: string
    score1: string
    score2: string
    round_info: string
    flag1: string
    flag2: string
    tournament_name: string
    tournament_icon: string
    chatId : string
}

interface getMatchsProps {
    setMatchs: React.Dispatch<React.SetStateAction<Match[]>>
}

export const getMatchs = async ({setMatchs}: getMatchsProps) => {
    const res = await axios.get(`${import.meta.env.VITE_BASE_URL}api/matchesResult`)
    const matchs = res.data.matchesResults
    setMatchs(matchs)
}