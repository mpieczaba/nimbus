import React from "react";
import { IconX } from "@tabler/icons";

import { colors } from "../../themes/colors";

import { Wrapper, TagName, Button } from "./styles";

interface Props {
  removable?: boolean;
  tagName: string;
  handleTagRemove?: () => void;
  onClick?: () => void;
}

const Tag: React.FC<Props> = ({
  removable,
  tagName,
  handleTagRemove,
  ...props
}) => {
  return (
    <Wrapper {...props}>
      <TagName>{tagName}</TagName>
      {removable ? (
        <Button onClick={handleTagRemove}>
          <IconX size="1rem" color={colors.text} />
        </Button>
      ) : null}
    </Wrapper>
  );
};

export default Tag;
