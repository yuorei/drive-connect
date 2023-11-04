import { useEffect, useState } from 'react'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import Typography from '@mui/material/Typography'
import { Button, CardActions, Box } from '@mui/material'

type Board = {
    id: string,
    type: string,
    user_id: string,
    description: string,
    departure_latitude: number,
    departure_longitude: number,
    destination_latitude: number,
    destination_longitude: number,
    reward: string,
    start_time: string,
    created_at: string,
    updated_at: string,

}

type Location = {
  longitude: number,
  latitude: number,
}

function Ride() {
  const [boards, setBoards] = useState<Board[]>([])
  const [location, setLocation] = useState<Location>({longitude: 0, latitude: 0})

  useEffect(() => {
    const newBoards = [
      {
        id: '1',
        type: 'drive',
        user_id: '1',
        description: 'test',
        departure_latitude: 0,
        departure_longitude: 0,
        destination_latitude: 0,
        destination_longitude: 0,
        reward: 'test',
        start_time: 'test',
        created_at: 'test',
        updated_at: 'test',
      },
      {
        id: '2',
        type: 'drive',
        user_id: '2',
        description: 'test',
        departure_latitude: 35,
        departure_longitude: 0,
        destination_latitude: 0,
        destination_longitude: 0,
        reward: 'test',
        start_time: 'test',
        created_at: 'test',
        updated_at: 'test',
      },
      {
        id: '3',
        type: 'drive',
        user_id: '3',
        description: 'test',
        departure_latitude: 0,
        departure_longitude: 0,
        destination_latitude: 0,
        destination_longitude: 0,
        reward: 'test',
        start_time: 'test',
        created_at: 'test',
        updated_at: 'test',
      }
    ]

    setBoards(newBoards.sort((a, b) => {
      return (location.latitude - a.departure_latitude) ** 2 + (location.longitude - a.departure_longitude) ** 2 >
      (location.latitude - b.departure_latitude) ** 2 + (location.longitude - b.departure_longitude) ** 2 ? 1 : -1
    }))
  }, [])


  useEffect(() => {
    navigator.geolocation.getCurrentPosition((p)=> {
      setLocation({
        longitude: p.coords.longitude,
        latitude: p.coords.latitude,
      })
    })
  }, [])
  
  return (
    <>
      <Box sx={{ margin: 10 }}>
        {boards.map((board) => <Card key={board.id} sx={{ width: 500 }}>
          <CardContent>
            <Typography color="text.secondary" align="left">{board.id}</Typography>
            <Typography>{board.description}</Typography>
            <Typography>reword: {board.description}</Typography>
            <Typography color="text.secondary" align="right" fontSize="1">最終更新日時: {board.updated_at}</Typography>
          </CardContent>
          <CardActions>
            <Button size="small">乗りたい</Button>
          </CardActions>
        </Card>)} 
      </Box>
    </>
  )
}

export default Ride
