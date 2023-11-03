import { useEffect, useState, createContext } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';

export const PositionContext = createContext({ latitude: 0, longitude: 0 });
export const MapContext = createContext({} as {
  map: google.maps.Map;
  setMap: React.Dispatch<React.SetStateAction<google.maps.Map>>;
});

function Register() {
  

  return (
    <>
      <Card>
        <CardContent
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <h1>Register</h1>
        </CardContent>
      </Card>
    </>
  )
}

export default Register
