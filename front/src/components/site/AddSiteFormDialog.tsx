import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React, { useEffect } from "react";
import {
  Box,
  TextField,
  DialogTitle,
  DialogContent,
  Button,
  Dialog,
  DialogContentText,
  DialogActions,
} from "@mui/material";

type FormInput = {
  name: string;
  url: string;
  feed_url: string;
};

const schema = yup.object({
  name: yup.string().required("必須です"),
  url: yup.string().required("必須です").url("正しいURLを入力してください"),
  feed_url: yup
    .string()
    .required("必須です")
    .url("正しいURLを入力してください"),
});

export type AddSiteFormProps = {
  open: boolean;
  handleClose: () => void;
  addSite: (name: string, url: string, feed_url: string) => void;
};

const AddSiteFormDialog: React.FC<AddSiteFormProps> = ({
  open,
  handleClose,
  addSite,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<FormInput>({
    defaultValues: {
      name: "",
      url: "",
      feed_url: "",
    },
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedName = data.name.trim();
      if (trimmedName.length === 0) return;

      const trimmedUrl = data.url.trim();
      if (trimmedUrl.length === 0) return;

      const trimmedfeed_url = data.feed_url.trim();
      if (trimmedfeed_url.length === 0) return;

      await addSite(trimmedName, trimmedUrl, trimmedfeed_url);
      reset();
      handleClose();
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    reset();
  }, [open]);

  return (
    <Dialog open={open} onClose={handleClose}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle>Add Site</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Siteの名前とURLを入力してください。
          </DialogContentText>
          <Box display="flex" flexDirection="column" gap={2}>
            <TextField
              label="Site Name"
              autoFocus={true}
              error={errors.name ? true : false}
              helperText={errors.name && errors.name.message}
              variant="outlined"
              placeholder="name"
              {...register("name")}
              sx={{ flexGrow: 1 }}
            />

            <TextField
              label="Site URL"
              error={errors.url ? true : false}
              helperText={errors.url && errors.url.message}
              variant="outlined"
              placeholder="url"
              {...register("url")}
              sx={{ flexGrow: 1 }}
            />

            <TextField
              label="Feed URL"
              error={errors.feed_url ? true : false}
              helperText={errors.feed_url && errors.feed_url.message}
              variant="outlined"
              placeholder="Feed Url"
              {...register("feed_url")}
              sx={{ flexGrow: 1 }}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button type="submit">Add</Button>
        </DialogActions>
      </Box>
    </Dialog>
  );
};

export default AddSiteFormDialog;
