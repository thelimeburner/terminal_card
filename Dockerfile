FROM arm64v8/alpine


WORKDIR /app

COPY terminal-card terminal_cmd

COPY card/. card/


ENV CC_KEY="duck"
ENV CC_SECRET="quack"
ENV CC_ADDRESS="54.89.51.42"
ENV CC_PORT="5000"

ENV CC_CARD_ENABLED=true

CMD ["/app/terminal_cmd"] 
