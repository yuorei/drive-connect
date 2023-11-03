import React, { useState } from 'react';
import { BottomNavigation, BottomNavigationAction, Paper } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import HailIcon from '@mui/icons-material/Hail';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';

export default () => {
  const [count, setCount] = useState(0);
  const navigate = useNavigate();
  const handleNavigation = (_event: React.SyntheticEvent, value: number) => {
    setCount(value);
    console.log(value);
    if (value === 0) {
      console.log('ride');
      return navigate('/ride');
    } else {
      console.log('drive');
      return navigate('/drive');
    }
  }
  return (
    <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
      <BottomNavigation
        showLabels
        value={count}
        onChange={handleNavigation}
      >
        <BottomNavigationAction label="行きたい！" icon={<HailIcon />} />
        <BottomNavigationAction label="乗せてもいいよ！" icon={<DirectionsCarIcon />} />
        </BottomNavigation>
    </Paper>
  )
}

