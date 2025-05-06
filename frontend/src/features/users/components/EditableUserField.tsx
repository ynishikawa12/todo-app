import React, { useEffect, useState } from "react";
import { Box, IconButton, TextField, Typography } from "@mui/material";
import { Edit } from "@mui/icons-material";

type EditableUserFieldProps = {
  label: string;
  fieldName: string;
  value: string;
};

export const EditableUserField = ({
  label,
  fieldName,
  value,
}: EditableUserFieldProps) => {
  const [isEditing, setIsEditing] = useState(false);
  const [inputValue, setInputValue] = useState<string>();

  useEffect(() => {
    setInputValue(value);
  }, [value]);

  const handleEditClick = () => {
    setIsEditing(true);
    setInputValue(value);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleSave = () => {
    setIsEditing(false);
  };

  const handleCancel = () => {
    setIsEditing(false);
    setInputValue(value);
  };

  const handleBlur = () => {
    handleSave();
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSave();
    } else if (e.key === "Escape") {
      handleCancel();
    }
  };

  const Editable = () => {
    return (
      <TextField
        autoFocus
        label={label}
        name={fieldName}
        variant="outlined"
        fullWidth
        margin="normal"
        value={inputValue}
        onChange={handleChange}
        onBlur={handleBlur}
        onKeyDown={handleKeyDown}
      />
    );
  };

  const Display = () => {
    return (
      <Box
        display="flex"
        alignItems="center"
        justifyContent="space-between"
        width={"100%"}
      >
        <Box flexGrow={1}>
          <Typography>{inputValue}</Typography>
        </Box>
        <IconButton onClick={handleEditClick}>
          <Edit />
        </IconButton>
      </Box>
    );
  };

  return (
    <Box display="flex" alignItems="center">
      {isEditing ? <Editable /> : <Display />}
    </Box>
  );
};
