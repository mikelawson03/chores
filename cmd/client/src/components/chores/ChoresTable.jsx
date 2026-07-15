import { Stack, Typography } from "@mui/material"
import TableHeader from "./TableHeader";
import TableToolbar from "./TableToolbar";
import ChoresTableRow from "./ChoresTableRow";
import { useState } from "react";
import dayjs from "dayjs";

export default function ChoresTable({ choreTemplates, openChoreDetails }) {
  const [sort, setSort] = useState(
      {
        column: "updated_at",
        direction: "desc"
      }
  )

  const filteredChores = [...choreTemplates];
  const displayedChores = sortChores(filteredChores, sort)

  
  function sortChores(chores, sort) {
    let comparison
    return chores.sort((a, b) => {
      switch (sort.column) {
        case "name":
          comparison = a.name.localeCompare(b.name);
          break;
        case "duration":
          comparison = a.duration - b.duration;
          break;
        case "updated_at":
          comparison = dayjs(a.updated_at).diff(dayjs(b.updated_at));
          break;
        
    }
    
    return sort.direction === "asc"
      ? comparison
      : -comparison
  })
  }

  function onSortClick(heading) {
    if (heading.field === sort.column) {
      setSort(previousSort => {
        return {
          ...previousSort,
          direction: previousSort.direction === "asc" ? "desc" : "asc"
        }
      })
    } else {
      setSort(previousSort => {
        return {
          column: heading.field,
          direction: heading.defaultDirection
        }
      })
    }
  }

  const tableHeadings=[
    {
      displayName: "Name",
      field: "name",
      sortable: true,
      defaultDirection: "asc",
    }, 
    {
      displayName: "Frequency",
      field: "cadence",
      sortable: false,
    }, 
    {
      displayName: "Default Assignee",
      field: "assignee",
      sortable: false,
    }, 
    {
      displayName: "Duration",
      field: "duration",
      sortable: true,
      defaultDirection: "asc",
    }, 
    {
      displayName: "Last Updated",
      field: "updated_at",
      sortable: true,
      defaultDirection: "desc",
    }
  ]
  return (
    <Stack direction="column" sx={{
        flex: 1, 
        width: "100%"
      }
    }>
      <TableToolbar />
      <TableHeader 
        sort={sort} 
        chores={displayedChores} 
        headings={tableHeadings} 
        onSortClick={onSortClick}
      />
      {displayedChores.map( chore => (
      <ChoresTableRow key={chore.id} choreTemplate={chore} openChoreDetails={openChoreDetails} />))}
    </Stack>
  )
}