import dayjs from "dayjs"

export function formatCadence(cadence) {
  const cadenceLabels = {
    daily: "Daily",
    weekly: "Weekly",
    monthly: "Monthly",
    annually: "Annually"
  }

  return cadenceLabels[cadence]
}

export function formatDuration(duration) {
  if (duration < 60) {
    return `${duration}m`
  }

  const hours = Math.floor(duration / 60)
  const mins = duration % 60

  if (mins === 0) {
    return `${hours}h`
  }

  return `${hours}h ${mins}m`
}

export function formatTimestamp(timestamp) {
  return dayjs(timestamp).format("MMM DD, YYYY • h:mm A")
}