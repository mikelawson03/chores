import { Box, Stack } from '@mui/material';
import ArrowDropUpIcon from '@mui/icons-material/ArrowDropUp';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import { inactiveIndicator } from '../styles/indicators';

export default function SortIndicator({ sort, heading }) {
    if (sort.column === heading.field) {
        if (sort.direction === "desc") {
            return <ArrowDropDownIcon />
        }
        return <ArrowDropUpIcon />
    }

    return <SwapVertIcon sx={inactiveIndicator} />
}