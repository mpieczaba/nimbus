import styled from "styled-components";
import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;

  svg {
    width: 1.5rem;
    height: 1.5rem;
  }
`;

export const Thumbnail = styled.img`
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;
