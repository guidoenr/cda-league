import React, { useState, useEffect } from 'react';
import Team from './Team'
import './Match.css'

const Match = () => {
    const [Team1, setTeam1] = useState({});
    const [Team2, setTeam2] = useState({});

    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/generateMatch/');
                const data = await response.json();
                setTeam1(data.team1);
                setTeam2(data.team2);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);

        return (
            <div className="match-container">
                <Team className="team-container"
                    name="TEAM 1"
                    players={Team1.players}
                    totalPlayers={Team1.totalPlayers}
                    chanceOfWinning={Team1.chanceOfWinning}
                />
                <Team className="team-container"
                    name="TEAM 2"
                    players={Team2.players}
                    totalPlayers={Team2.totalPlayers}
                    chanceOfWinning={Team2.chanceOfWinning}
                />
            </div>
        );
}



export default Match;
