import { Alert, Box } from "@mui/material";

interface AlertMessageProps {
    message: string;
}

export const AlertMessage = ({ message }: AlertMessageProps) => {
    return (
        message &&
        <Box sx={{ mt: 2 }}>
            <Alert severity={message.includes('successfully') ? 'success' : 'error'}>
                {message}
            </Alert>
        </Box>
    );
}