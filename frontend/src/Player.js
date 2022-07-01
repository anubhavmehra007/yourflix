import React from "react";
import {useParams} from 'react-router-dom';
export default function Player() {
    const {id} = useParams();
    const [movie, setMovie] = React.useState({});
    React.useEffect(() => {
    return async () =>  {
        const fetchOpts = {
            method: 'GET',
            mode: 'cors'
        };
        const url = `http://localhost:8000/movie/${id}`;
        console.log(url);
        const response = await fetch(url, fetchOpts);
        setMovie(await response.json());
    }},[id]);
    return (
        <div>
            <h1>{movie.Name}</h1>
            <h2>{movie.Director}</h2>
            <video  controls key={movie.Id}>
                <source src={movie.Id ? `http://localhost:8000/play/${movie.Id}` : "deafult.mp4"} type="video/mp4"></source>
            </video>
        </div>
    )

}