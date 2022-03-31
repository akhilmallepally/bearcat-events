import { useState, useEffect, Fragment } from "react"

const List = () => {
  const [eventsList, setEventsList] = useState([])

  useEffect(() => {
    fetch(process.env.BEARCATEVENTS_API)
      .then((res) => res.json())
      .then(setEventsList)
  }, [])

  if (eventsList.length !== 0) {
    console.log(eventsList)
    return (
      
      <Fragment>
      
      <div>
        {eventsList.map((data, index) => {
          return (
            <><div className="card col-sm-3">
              <div className="container">
                <div className="card-body">
                  <h5 key={index} className="card-title" color="blue"><strong>{data.title}</strong></h5>
                  <p key={index} className="card-text">{data.description}</p>
                </div>
                <div className="card-footer">
                  <small className="text-muted">Last updated 3 mins ago</small>
                </div>
              </div>
            </div>
              </>
          )
        })}
      </div>

      </Fragment>)
    
  } else {
    return <h3>Loading...</h3>
  }
}

export default List
