import React, { useState, useEffect } from 'react';
import Team from './Team'
import './Match.css'

const Match = () => {
    const [Team1, setTeam1] = useState({});
    const [Team2, setTeam2] = useState({});

    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/match');
                const data = await response.json();
                console.log("DATA" + data)
                console.log("WTF")
                setTeam1({...data.team1});
                setTeam2({...data.team2});
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);


    return (
        <div className="match-container">
            <Team {...Team1} />
            <Team {...Team2} />
        </div>
    );
}

export default Match;
