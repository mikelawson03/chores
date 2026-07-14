export function formatCadence(cadence) {
  const cadenceLabels = {
    daily: "Daily",
    weekly: "Weekly",
    monthly: "Monthly",
    annually: "Annually"
  }

  return cadenceLabels[cadence]
}