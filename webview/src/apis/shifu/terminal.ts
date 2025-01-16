import {
  CommandRequest,
  CommandResponse,
  CompletionRequest,
  CompletionResponse,
} from "src/proto/proto/shifu/shifu_pb";
import { client } from "./shifu";
import { ClientReadableStream } from "grpc-web";

export function ExecuteCommand(command: string): {
  stream: ClientReadableStream<CommandResponse>;
  cancel: () => void;
} {
  const request = new CommandRequest();
  request.setCommand(command);
  const stream = client.executeCommand(request);
  return {
    stream,
    cancel: () => {
      stream.cancel();
    },
  };
}

export const GetCompletions = async (
  partial: string,
  currentDir: string
): Promise<string[]> => {
  try {
    const request = new CompletionRequest();
    request.setPartial(partial);
    request.setCurrentDir(currentDir);

    return new Promise((resolve, reject) => {
      client.getCompletions(request, {}, (err, response) => {
        if (err) {
          console.error("Failed to get completions:", err);
          reject(err);
          return;
        }
        resolve(response.getCompletionsList());
      });
    });
  } catch (error) {
    console.error("Error getting completions:", error);
    return [];
  }
};
