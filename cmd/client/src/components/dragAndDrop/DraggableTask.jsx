import {useDraggable} from '@dnd-kit/react';

export default function DraggableTask({ task, children }) {
    const { ref, isDragging } = useDraggable({
        id: task.id
    });

    return (
        <div ref={ref}>
            {children(isDragging)}
        </div>
    )
}