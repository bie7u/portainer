import { PropsWithChildren, ReactNode } from 'react';
import { SchemaOf, string } from 'yup';

import { StackId } from '@/react/common/stacks/types';

import { FormControl } from '@@/form-components/FormControl';
import { TextTip } from '@@/Tip/TextTip';

import { RefSelector } from './RefSelector';
import { RefFieldModel } from './types';

interface Props {
  value: string;
  onChange(value: string): void;
  model: RefFieldModel;
  error?: string;
  isUrlValid?: boolean;
  stackId?: StackId;
  createdFromCustomTemplateId?: number;
}

export function RefField({
  value,
  onChange,
  model,
  error,
  isUrlValid,
  stackId,
  createdFromCustomTemplateId,
}: Props) {
  const inputId = 'repository-reference-field';
  return (
    <Wrapper
      inputId={inputId}
      errors={error}
      tip={
        <>
          Specify a reference of the repository using the following syntax:
          branches with <code>refs/heads/branch_name</code> or tags with{' '}
          <code>refs/tags/tag_name</code>.
        </>
      }
    >
      <RefSelector
        inputId={inputId}
        value={value}
        onChange={onChange}
        model={model}
        isUrlValid={isUrlValid}
        stackId={stackId}
        createdFromCustomTemplateId={createdFromCustomTemplateId}
      />
    </Wrapper>
  );
}

function Wrapper({
  tip,
  children,
  errors,
  inputId,
}: PropsWithChildren<{ tip: ReactNode; errors?: string; inputId: string }>) {
  return (
    <div className="form-group">
      <span className="col-sm-12 mb-2">
        <TextTip color="blue">{tip}</TextTip>
      </span>
      <div className="col-sm-12">
        <FormControl
          label="Repository reference"
          inputId={inputId}
          required
          errors={errors}
        >
          {children}
        </FormControl>
      </div>
    </div>
  );
}

export function refFieldValidation(): SchemaOf<string> {
  return string().required('Repository reference name is required').default('');
}
